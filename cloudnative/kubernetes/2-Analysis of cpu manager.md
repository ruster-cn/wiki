# 2-CPU Manager 源码分析

## CPU Manager 是什么

默认情况下，kubelet 创建的 pod 都是使用 CFS 配额的方式使用 CPU 资源的。在这种分配方式下，所有的容器可以使用物理机所有的 cpu 核，并根据申请的 cpu 资源的比例抢占 cpu 时钟周期。CFS 配额的方式对于偏计算的服务，不是特别友好。主要在于进程每次抢占到 cpu 资源后，都会存在上下文切换的问题，以及高速缓存丢失等问题。

cpu manager 提供了一种 cpu set 的 cpu 分配方式。可以将进程尽可能的绑定到同一个 numa 节点上的几个核上运行，减少运行期间的上下文切换。因此如果服务对 cpu throttling、上下文切换、处理器高速缓存丢失、跨插槽的内存访问、需要使用统一物理核的超线程 cpu 敏感的服务，都应该使用 cpu manager。

关于 cpu manager 到底能带来多大的性能提升，可以参考官方给出的一个性能测试结果。![](https://d33wubrfki0l68.cloudfront.net/722e38cc2f45525d9904e6103b44ad221a469df1/9e871/images/blog/2018-07-24-cpu-manager/execution-time.png)

## 代码剖析

cpu manager 提供了两种策略：none 和 static。kubelet 默认的 cpu-manager-policy 是 none，即使用 CFS 配额的方式分配使用 cpu 资源。如果将 static策略针对具有整数型cpu request的Guananteed Pod，允许该类Pod中的容器访问节点上独占的cpu资源。static policy管理一个共享的cpu资源池。最初，该资源池包含节点上所有的cpu资源。可以独占的cpu资源等于节点的 cpu总量 - kube-reserved保留的cpu - system-reserved保留的cpu.从1.17版本后，cpu保留列表可以通过reserved-cpus参数显示的设置。reserved-cpus指定的cpu列表优先于kube-reserved 和 system-reserved。 通过这些参数预留的 CPU 是以整数方式，按物理内 核 ID 升序从初始共享池获取的。 共享池是 BestEffort 和 Burstable pod 运行 的 CPU 集合。Guaranteed pod 中的容器，如果声明了非整数值的 CPU requests ，也将运行在共享池的 CPU 上。只有 Guaranteed pod 中，指定了整数型 CPU requests 的容器，才会被分配独占 CPU 资源。
在下面的代码分析中将会主要分析 static cpu manager policy.因为在 none cpu manager policy 的代码实现都是空实现。

```golang
type Manager interface {
    //kubelet 启动时调用启动cpu manager
    Start(activePods ActivePodsFunc, sourcesReady config.SourcesReady, podStatusProvider status.PodStatusProvider, containerRuntime runtimeService, initialContainers containermap.ContainerMap) error
    //kubelet 做准入控制时调用，检查目前机器上cpu资源是否足够分配给pod。在1.17 版本之前，cpu manager 没有定时的将内存和机器cpu set资源做同步。当用户手动清理了容器后，分配给容器的cpu set资源不会被回收。但是机器的cpu 资源是充足的。因此pod可以被调度到机器上，但是kubelet做准入检查时，就会失败,产生PreStartHookError。
    Allocate(pod *v1.Pod, container *v1.Container) error
    //调用runtime运行容器前调用，更新容器cpu set设置。
    AddContainer(p *v1.Pod, c *v1.Container, containerID string) error
    //容器销毁时调用，回收分配出去的cpu set资源。
    RemoveContainer(containerID string) error
    //获取目前机器cpu set分配情况
    State() state.Reader
    //用于和其他模块同步numa拓扑
	GetTopologyHints(*v1.Pod, *v1.Container) map[string][]topologymanager.TopologyHint
}
```

### 1. 初始化

```golang
// NewManager creates new cpu manager based on provided policy
func NewManager(cpuPolicyName string, reconcilePeriod time.Duration, machineInfo *cadvisorapi.MachineInfo, numaNodeInfo topology.NUMANodeInfo, specificCPUs cpuset.CPUSet, nodeAllocatableReservation v1.ResourceList, stateFileDirectory string, affinity topologymanager.Store) (Manager, error) {
...
	switch policyName(cpuPolicyName) {
	case PolicyNone:
		policy = NewNonePolicy()
	case PolicyStatic:
		var err error
		topo, err = topology.Discover(machineInfo, numaNodeInfo)
		if err != nil {
			return nil, err
		}
		klog.Infof("[cpumanager] detected CPU topology: %v", topo)
		reservedCPUs, ok := nodeAllocatableReservation[v1.ResourceCPU]
...
		policy, err = NewStaticPolicy(topo, numReservedCPUs, specificCPUs, affinity)

	}
	manager := &manager{
		policy:                     policy,
		reconcilePeriod:            reconcilePeriod,
		topology:                   topo,
		nodeAllocatableReservation: nodeAllocatableReservation,
		stateFileDirectory:         stateFileDirectory,
	}
	manager.sourcesReady = &sourcesReadyStub{}
	return manager, nil
}
```

- 首先 manager 会根据 cpuPolicyName 生成不同的 policy，none policy 基本没有实现，可以认为什么都不做。

### 2. 启动

### 3. 准入检查

### 4. 添加容器

### 5. 删除容器



## cpu manager 与 topology manager && device manager
