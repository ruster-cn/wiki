# Cgroup中的CPU资源控制

在/sys/fs/cgroup下面,可以看到三个和CPU相关的子系统:cpu,cpuacct,cpuset.

## cpu subsystem
cpu子系统用于控制cgroup中所有进程可以使用的cpu时间片。

cpu subsystem主要涉及5接口:cpu.cfs_period_us，cpu.cfs_quota_us，cpu.shares，cpu.rt_period_us，cpu.rt_runtime_us.

#### cpu.cfs_period_us
cfs_period_us表示一个cpu带宽，单位为微秒。系统总CPU带宽： cpu核心数 * cfs_period_us

#### cpu.cfs_quota_us
cfs_quota_us表示Cgroup可以使用的cpu的带宽，单位为微秒。cfs_quota_us为-1，表示使用的CPU不受cgroup限制。cfs_quota_us的最小值为1ms(1000)，最大值为1s。

结合cfs_period_us,就可以限制进程使用的cpu。例如配置cfs_period_us=10000，而cfs_quota_us=2000。那么该进程就可以可以用2个cpu core。

#### cpu.shares
通过cfs_period_us和cfs_quota_us可以以绝对比例限制cgroup的cpu使用，即cfs_quota_us/cfs_period_us 等于进程可以利用的cpu cores，不能超过这个数值。

而cpu.shares以相对比例限制cgroup的cpu。例​​​如​​​：在​​​两​​​个​​​ cgroup 中​​​都​​​将​​​ cpu.shares 设​​​定​​​为​​​ 1 的​​​任​​​务​​​将​​​有​​​相​​​同​​​的​​​ CPU 时​​​间​​​，但​​​在​​​ cgroup 中​​​将​​​ cpu.shares 设​​​定​​​为​​​ 2 的​​​任​​​务​​​可​​​使​​​用​​​的​​​ CPU 时​​​间​​​是​​​在​​​ cgroup 中​​​将​​​ cpu.shares 设​​​定​​​为​​​ 1 的​​​任​​​务​​​可​​​使​​​用​​​的​​​ CPU 时​​​间​​​的​​​两​​​倍​​​。​​​

#### cpu.rt_runtime_us
以​​​微​​​秒​​​（µs，这​​​里​​​以​​​“​​​us”​​​代​​​表​​​）为​​​单​​​位​​​指​​​定​​​在​​​某​​​个​​​时​​​间​​​段​​​中​​​ cgroup 中​​​的​​​任​​​务​​​对​​​ CPU 资​​​源​​​的​​​最​​​长​​​连​​​续​​​访​​​问​​​时​​​间​​​。​​​建​​​立​​​这​​​个​​​限​​​制​​​是​​​为​​​了​​​防​​​止​​​一​​​个​​​ cgroup 中​​​的​​​任​​​务​​​独​​​占​​​ CPU 时​​​间​​​。​​​如​​​果​​​ cgroup 中​​​的​​​任​​​务​​​应​​​该​​​可​​​以​​​每​​​ 5 秒​​​中​​​可​​​有​​​ 4 秒​​​时​​​间​​​访​​​问​​​ CPU 资​​​源​​​，请​​​将​​​ cpu.rt_runtime_us 设​​​定​​​为​​​ 4000000，并​​​将​​​ cpu.rt_period_us 设​​​定​​​为​​​ 5000000。​​​

#### cpu.rt_period_us
以​​​微​​​秒​​​（µs，这​​​里​​​以​​​“​​​us”​​​代​​​表​​​）为​​​单​​​位​​​指​​​定​​​在​​​某​​​个​​​时​​​间​​​段​​​中​​​ cgroup 对​​​ CPU 资​​​源​​​访​​​问​​​重​​​新​​​分​​​配​​​的​​​频​​​率​​​。​​​如​​​果​​​某​​​个​​​ cgroup 中​​​的​​​任​​​务​​​应​​​该​​​每​​​ 5 秒​​​钟​​​有​​​ 4 秒​​​时​​​间​​​可​​​访​​​问​​​ CPU 资​​​源​​​，则​​​请​​​将​​​ cpu.rt_runtime_us 设​​​定​​​为​​​ 4000000，并​​​将​​​ cpu.rt_period_us 设​​​定​​​为​​​ 5000000。​​​

下面通过kubernetes创建一个burstable类型的,观察cpu subsystem各项的参数。

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    component: nginx
  name: nginx
spec:
  containers:
      image: nginx
      imagePullPolicy: IfNotPresent
      name: nginx
      resources:
        requests:
          cpu: 1
          memory: 1Gi
```

```shell
# cd /sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-podd8b74b92f7c532eac09641968762d50c.slice
# cd cri-containerd-82e854d26fe65d318f99819dd7e960f12d91ffccc48965b9d86d3b842e11dad7.scope
# ll
total 0
-rw-r--r-- 1 root root 0 Jan 21 13:22 cgroup.clone_children
--w--w--w- 1 root root 0 Jan 21 13:22 cgroup.event_control
-rw-r--r-- 1 root root 0 Jan 21 13:22 cgroup.procs
-r--r--r-- 1 root root 0 Jan 21 13:22 cpuacct.stat
-rw-r--r-- 1 root root 0 Jan 21 13:22 cpuacct.usage
-r--r--r-- 1 root root 0 Jan 21 13:22 cpuacct.usage_percpu
-rw-r--r-- 1 root root 0 Jan 21 18:34 cpu.cfs_period_us
-rw-r--r-- 1 root root 0 Jan 21 13:22 cpu.cfs_quota_us
-rw-r--r-- 1 root root 0 Jan 21 13:22 cpu.rt_period_us
-rw-r--r-- 1 root root 0 Jan 21 13:22 cpu.rt_runtime_us
-rw-r--r-- 1 root root 0 Jan 21 18:34 cpu.shares
-r--r--r-- 1 root root 0 Jan 21 13:22 cpu.stat
-rw-r--r-- 1 root root 0 Jan 21 13:22 notify_on_release
-rw-r--r-- 1 root root 0 Jan 21 13:22 tasks

# cat cpu.cfs_period_us
100000

# cat cpu.cfs_quota_us
-1

# cat cpu.shares
1024
```

如上所示，该container没有使用绝对方式控制cpu 资源使用，而是使用cpu share的方式控制cpu资源的访问。

## cpuacct subsystem
cpuacct子系统（CPU accounting）会自动生成报告来显示cgroup中任务所使用的CPU资源。报告有两大类：
cpuacct.stat和cpuacct.usage。


#### cpuacct.stat
cpuacct.stat记录cgroup的所有任务（包括其子孙层级中的所有任务）使用的用户和系统CPU时间.

```shell
# cat cpuacct.stat
user 2030466 #用户模式中任务使用的CPU时间
system 623280 #系统模式(内核)中任务使用的CPU时间
```
#### cpuacct.usage
cpuacct.usage记录这​​​个​​​cgroup中​​​所​​​有​​​任​​​务​​​（包括其子孙层级中的所有任务）消​​​耗​​​的​​​总​​​CPU时​​​间​​​（纳​​​秒​​​）。

#### cpuacct.usage_percpu
cpuacct.usage_percpu记录​​​这​​​个​​​cgroup中​​​所​​​有​​​任​​​务​​​（包括其子孙层级中的所有任务​​​）在​​​每​​​个​​​CPU中​​​消​​​耗​​​的​​​CPU时​​​间​​​（以​​​纳​​​秒​​​为​​​单​​​位​​​）。​​​

## cpuset subsystem
cpuset主要是为了numa使用的，numa技术将CPU划分成不同的node，每个node由多个CPU组成，并且有独立的本地内存、I/O等资源(硬件上保证)。可以使用numactl查看当前系统的node信息。


#### cpuset.cpus
cpuset.cpus指​​​定​​​允​​​许​​​这​​​个​​​ cgroup 中​​​任​​​务​​​访​​​问​​​的​​​ CPU。​​​这​​​是​​​一​​​个​​​用​​​逗​​​号​​​分​​​开​​​的​​​列​​​表​​​，格​​​式​​​为​​​ ASCII，使​​​用​​​小​​​横​​​线​​​（"-"）代​​​表​​​范​​​围​​​。​​​如下，代​​​表​​​ CPU 0、​​​1、​​​2 和​​​ 16。​​​
```shell
# cat cpuset.cpus
0-2,16
```

#### cpuset.mems
cpuset.mems指​​​定​​​允​​​许​​​这​​​个​​​ cgroup 中​​​任​​​务​​​可​​​访​​​问​​​的​​​内​​​存​​​节​​​点​​​。​​​这​​​是​​​一​​​个​​​用​​​逗​​​号​​​分​​​开​​​的​​​列​​​表​​​，格​​​式​​​为​​​ ASCII，使​​​用​​​小​​​横​​​线​​​（"-"）代​​​表​​​范​​​围​​​。​​​如下代​​​表​​​内​​​存​​​节​​​点​​​ 0、​​​1、​​​2 和​​​ 16。​​​

```shell
# cat cpuset.mems
0-2,16
```

#### cpuset.memory_migrate
cpuset.memory_migrate 用​​​来​​​指​​​定​​​当​​​ cpuset.mems 中​​​的​​​值​​​更​​​改​​​时​​​是​​​否​​​应​​​该​​​将​​​内​​​存​​​中​​​的​​​页​​​迁​​​移​​​到​​​新​​​节​​​点​​​的​​​标​​​签​​​（0 或​​​者​​​ 1）。​​​默​​​认​​​情​​​况​​​下​​​禁​​​止​​​内​​​存​​​迁​​​移​​​（0）且​​​页​​​就​​​保​​​留​​​在​​​原​​​来​​​分​​​配​​​的​​​节​​​点​​​中​​​，即​​​使​​​在​​​ cpuset.mems 中​​​现​​​已​​​不​​​再​​​指​​​定​​​这​​​个​​​节​​​点​​​。​​​如​​​果​​​启​​​用​​​（1），则​​​该​​​系​​​统​​​会​​​将​​​页​​​迁​​​移​​​到​​​由​​​ cpuset.mems 指​​​定​​​的​​​新​​​参​​​数​​​中​​​的​​​内​​​存​​​节​​​点​​​中​​​，可​​​能​​​的​​​情​​​况​​​下​​​保​​​留​​​其​​​相​​​对​​​位​​​置​​​ - 例​​​如​​​：原​​​来​​​由​​​ cpuset.mems 指​​​定​​​的​​​列​​​表​​​中​​​第​​​二​​​个​​​节​​​点​​​中​​​的​​​页​​​将​​​会​​​重​​​新​​​分​​​配​​​给​​​现​​​在​​​由​​​ cpuset.mems 指​​​定​​​的​​​列​​​表​​​的​​​第​​​二​​​个​​​节​​​点​​​中​​​，如​​​果​​​这​​​个​​​位​​​置​​​是​​​可​​​用​​​的​​​。​​​

#### cpuset.cpu_exclusive
puset.cpu_exclusive 指​​​定​​​​​​其​​​它​​​ cpuset 及​​​其​​​上​​​、​​​下​​​级cgroup是​​​否可​​​共​​​享​​​为​​​这​​​个​​​ cpuset 指​​​定​​​的​​​ CPU 的​​​标​​​签​​​（0 或​​​者​​​ 1）。​​​默​​​认​​​情​​​况​​​下​​​（0）CPU 不​​​是​​​专​​​门​​​分​​​配​​​给​​​某​​​个​​​ cpuset 的​​​。​​​

#### cpuset.mem_exclusive
cpuset.mem_exclusive 指​​​定​​​​​​其​​​它​​​ cpuset 是​​​否可​​​共​​​享​​​为​​​这​​​个​​​ cpuset 指​​​定​​​的​​​内​​​存​​​节​​​点​​​的​​​标​​​签​​​（0 或​​​者​​​ 1）。​​​默​​​认​​​情​​​况​​​下​​​（0）内​​​存​​​节​​​点​​​不​​​是​​​专​​​门​​​分​​​配​​​给​​​某​​​个​​​ cpuset 的​​​。​​​专​​​门​​​为​​​某​​​个​​​ cpuset 保​​​留​​​内​​​存​​​节​​​点​​​（1）与​​​使​​​用​​​ cpuset.mem_hardwall 启​​​用​​​内​​​存​​​ hardwall 功​​​能​​​是​​​一​​​致​​​的​​​。​​​

#### cpuset.mem_hardwall
cpuset.mem_hardwall指​​​定​​​是​​​否​​​应​​​将​​​内​​​存​​​页​​​面​​​的​​​内​​​核​​​分​​​配​​​限​​​制​​​在​​​为​​​这​​​个​​​ cpuset 指​​​定​​​的​​​内​​​存​​​节​​​点​​​的​​​标​​​签​​​（0 或​​​者​​​ 1）。​​​默​​​认​​​情​​​况​​​下​​​为​​​ 0，属​​​于​​​多​​​个​​​用​​​户​​​的​​​进​​​程​​​共​​​享​​​页​​​面​​​和​​​缓​​​冲​​​。​​​启​​​用​​​ hardwall 时​​​（1）每​​​个​​​任​​​务​​​的​​​用​​​户​​​分​​​配​​​应​​​保​​​持​​​独​​​立​​​。​​​

#### cpuset.memory_pressure
cpuset.memory_pressure是运​​​行​​​在​​​这​​​个​​​ cpuset 中​​​产​​​生​​​的​​​平​​​均​​​内​​​存​​​压​​​力​​​的​​​只​​​读​​​文​​​件​​​。​​​启​​​用​​​ cpuset.memory_pressure_enabled 时​​​，这​​​个​​​伪​​​文​​​件​​​中​​​的​​​值​​​会​​​自​​​动​​​更​​​新​​​，否​​​则​​​伪​​​文​​​件​​​包​​​含​​​的​​​值​​​为​​​ 0。​​​

#### cpuset.memory_pressure_enabled
cpuset.memory_pressure_enabled指​​​定​​​系​​​统​​​是​​​否​​​应​​​该​​​计​​​算​​​这​​​个​​​ cgroup 中​​​进​​​程​​​所​​​生​​​成​​​内​​​存​​​压​​​力​​​的​​​标​​​签​​​（0 或​​​者​​​ 1）。​​​计​​​算​​​出​​​的​​​值​​​会​​​输​​​出​​​到​​​ cpuset.memory_pressure，且​​​代​​​表​​​进​​​程​​​试​​​图​​​释​​​放​​​使​​​用​​​中​​​内​​​存​​​的​​​比​​​例​​​，报​​​告​​​为​​​尝​​​试​​​每​​​秒​​​再​​​生​​​内​​​存​​​的​​​整​​​数​​​值​​​再​​​乘​​​ 1000。​​​

#### cpuset.memory_spread_page
cpuset.memory_spread_page指​​​定​​​是​​​否​​​应​​​将​​​文​​​件​​​系​​​统​​​缓​​​冲​​​平​​​均​​​分​​​配​​​给​​​这​​​个​​​ cpuset 的​​​内​​​存​​​节​​​点​​​的​​​标​​​签​​​（0 或​​​者​​​ 1）。​​​默​​​认​​​情​​​况​​​为​​​ 0，不​​​尝​​​试​​​为​​​这​​​些​​​缓​​​冲​​​平​​​均​​​分​​​配​​​内​​​存​​​页​​​面​​​，且​​​将​​​缓​​​冲​​​放​​​置​​​在​​​运​​​行​​​生​​​成​​​缓​​​冲​​​的​​​进​​​程​​​的​​​同​​​一​​​节​​​点​​​中​​​。​​​

#### cpuset.memory_spread_slab
cpuset.memory_spread_slab指​​​定​​​是​​​否​​​应​​​在​​​ cpuset 间​​​平​​​均​​​分​​​配​​​用​​​于​​​文​​​件​​​输​​​入​​​/输​​​出​​​操​​​作​​​的​​​内​​​核​​​缓​​​存​​​板​​​的​​​标​​​签​​​（0 或​​​者​​​ 1）。​​​默​​​认​​​情​​​况​​​是​​​ 0，即​​​不​​​尝​​​试​​​平​​​均​​​分​​​配​​​内​​​核​​​缓​​​存​​​板​​​，并​​​将​​​缓​​​存​​​板​​​放​​​在​​​生​​​成​​​这​​​些​​​缓​​​存​​​的​​​进​​​程​​​所​​​运​​​行​​​的​​​同​​​一​​​节​​​点​​​中​​​。​​​
##### cpuset.sched_load_balance
cpuset.sched_load_balance指​​​定​​​是​​​否​​​在​​​这​​​个​​​ cpuset 中​​​跨​​​ CPU 平​​​衡​​​负​​​载​​​内​​​核​​​的​​​标​​​签​​​（0 或​​​者​​​ 1）。​​​默​​​认​​​情​​​况​​​是​​​ 1，即​​​内​​​核​​​将​​​超​​​载​​​ CPU 中​​​的​​​进​​​程​​​移​​​动​​​到​​​负​​​载​​​较​​​低​​​的​​​ CPU 中​​​以​​​便​​​平​​​衡​​​负​​​载​​​。​​​
请​​​注​​​意​​​：如​​​果​​​在​​​任​​​意​​​上​​​级​​​ cgroup 中​​​启​​​用​​​负​​​载​​​平​​​衡​​​，则​​​在​​​ cgroup 中​​​设​​​定​​​这​​​个​​​标​​​签​​​没​​​有​​​任​​​何​​​效​​​果​​​，因​​​为​​​已​​​经​​​在​​​较​​​高​​​一​​​级​​​ cgroup 中​​​处​​​理​​​了​​​负​​​载​​​平​​​衡​​​。​​​因​​​此​​​，要​​​在​​​ cgroup 中​​​禁​​​用​​​负​​​载​​​平​​​衡​​​，还​​​要​​​在​​​该​​​层​​​级​​​的​​​每​​​一​​​个​​​上​​​级​​​ cgroup 中​​​禁​​​用​​​负​​​载​​​平​​​衡​​​。​​​这​​​里​​​您​​​还​​​应​​​该​​​考​​​虑​​​是​​​否​​​应​​​在​​​所​​​有​​​平​​​级​​​ cgroup 中​​​启​​​用​​​负​​​载​​​平​​​衡​​​。​​​
#### cpuset.sched_relax_domain_level
cpuset.sched_relax_domain_level包​​​含​​​ -1 到​​​小​​​正​​​数​​​间​​​的​​​整​​​数​​​，它​​​代​​​表​​​内​​​核​​​应​​​尝​​​试​​​平​​​衡​​​负​​​载​​​的​​​ CPU 宽​​​度​​​范​​​围​​​。​​​如​​​果​​​禁​​​用​​​了​​​ cpuset.sched_load_balance，则​​​该​​​值​​​毫​​​无​​​意​​​义​​​。


下面通过kubernetes创建一个Guaranteed类型的,同时kubelet支持static cpu policy,观察cpu subsystem各项的参数。

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    component: nginx
  name: nginx
spec:
  containers:
      image: nginx
      imagePullPolicy: IfNotPresent
      name: nginx
      resources:
        limits: 
          cpu: "1"
          memory: "1Gi"
        requests:
          cpu: "1"
          memory: "1Gi"
```

```shell
# cat cpu.cfs_period_us
100000

# cat cpu.cfs_quota_us
100000

# cat cpu.shares
1024

# cat cpuset.cpus
32

# cat cpuset.mems
0-1
```

从cgrop配置中可以看到，container只占用了32 core。cpu.cfs_quota_us = cpu.cfs_period_us，容器可以完全使用整个cpu core。

如下是测试机numa节点分布：
```
# numactl -H
available: 2 nodes (0-1)
node 0 cpus: 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47
node 0 size: 65209 MB
node 0 free: 55299 MB
node 1 cpus: 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63
node 1 size: 65536 MB
node 1 free: 55112 MB
node distances:
node   0   1
  0:  10  21
  1:  21  10
```