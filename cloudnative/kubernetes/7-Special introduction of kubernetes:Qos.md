# kubernetes 专题介绍(1): Qos

k8s Pod的Qos类型包括以下三种:
* Guaranteed
* Burstable
* BestEffort

## Pod Qos 判定
具体的判定Pod Qos的算法实现可以参考如下代码：
```golang
// pkg/apis/core/v1/helper/qos/qos.go
/ GetPodQOS returns the QoS class of a pod.
// A pod is besteffort if none of its containers have specified any requests or limits.
// A pod is guaranteed only when requests and limits are specified for all the containers and they are equal.
// A pod is burstable if limits and requests do not match across all containers.
func GetPodQOS(pod *v1.Pod) v1.PodQOSClass {
	requests := v1.ResourceList{}
	limits := v1.ResourceList{}
	zeroQuantity := resource.MustParse("0")
	isGuaranteed := true
	allContainers := []v1.Container{}
	allContainers = append(allContainers, pod.Spec.Containers...)
	allContainers = append(allContainers, pod.Spec.InitContainers...)
	for _, container := range allContainers {
		// process requests
		for name, quantity := range container.Resources.Requests {
			if !isSupportedQoSComputeResource(name) {
				continue
			}
			if quantity.Cmp(zeroQuantity) == 1 {
				delta := quantity.DeepCopy()
				if _, exists := requests[name]; !exists {
					requests[name] = delta
				} else {
					delta.Add(requests[name])
					requests[name] = delta
				}
			}
		}
		// process limits
		qosLimitsFound := sets.NewString()
		for name, quantity := range container.Resources.Limits {
			if !isSupportedQoSComputeResource(name) {
				continue
			}
			if quantity.Cmp(zeroQuantity) == 1 {
				qosLimitsFound.Insert(string(name))
				delta := quantity.DeepCopy()
				if _, exists := limits[name]; !exists {
					limits[name] = delta
				} else {
					delta.Add(limits[name])
					limits[name] = delta
				}
			}
		}

		if !qosLimitsFound.HasAll(string(v1.ResourceMemory), string(v1.ResourceCPU)) {
			isGuaranteed = false
		}
	}
	if len(requests) == 0 && len(limits) == 0 {
		return v1.PodQOSBestEffort
	}
	// Check is requests match limits for all resources.
	if isGuaranteed {
		for name, req := range requests {
			if lim, exists := limits[name]; !exists || lim.Cmp(req) != 0 {
				isGuaranteed = false
				break
			}
		}
	}
	if isGuaranteed &&
		len(requests) == len(limits) {
		return v1.PodQOSGuaranteed
	}
	return v1.PodQOSBurstable
}

```

从以上代码可以了解,计算Qos只关心pod的cpu 和 mem 资源的申请情况：
* Guaranteed：
  
  * Pod limit中cpu和mem都不必须不等于0
  * Pod cpu和mem的limit和request的**总量（所有container的和，和每一contaienr无关，只要能保证最后Pod中所有container的和是相等即可**必须相等

* BestEffort:  
  
  * Pod request 中的cpu和mem 都等于0 或者 limit中的cpu和mem都等于0

* Burstable: 除去以上2种的


## Qos的应用

* cpu manager policy：如果cpu 

```
// pkg/kubelet/cm/cpumanager/policy_static.go#guaranteedCPUs
```

* 计算container oom score:

```
//pkg/kubelet/qos/policy.go#GetContainerOOMScoreAdjust
    //container oom score 计算规则:
    if pod is Guaranteed:
        return -998
    else if pod is BestEffort:
        return 1000
    else:
        oomScoreAdjust = 1000 - (1000*memoryRequest)/memoryCapacity
        if oomScoreAdjust < 2:
            return 2
        if oomScoreAdjust = 1000:
            return 999 //(oomScoreAdjust-1)
```
* container ResourceConfig 计算:

```
// pkg/kubelet/cm/qos_container_manager_linux.go#setCPUCgroupConfig
// pkg/kubelet/cm/qos_container_manager_linux.go#setMemoryReserve
```

* container cgroup path 设置:
  https://www.cnblogs.com/sparkdev/p/9523194.html
  https://blog.tianfeiyu.com/2020/01/21/kubelet_qos/
  
* pod 驱逐与抢占:
  
```
// pkg/kubelet/eviction/eviction_manager.go#Admit
// pkg/kubelet/preemption/preemption.go#evictPodsToFreeRequests
```  