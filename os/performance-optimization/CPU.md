# CPU 性能优化学习笔记

## 平均负载

> 定义：平均负责是指单位时间内，系统处于运行和不可中断状态的平均进程数，也称为平均活跃进程数。

```
linux上进程常见5种状态：
* 运行状态：进程正在运行(即正在使用cpu)或者在运行队列种等待(即正在等待cpu)
    ps 工具表示的状态码： R running or runnable (on run queue)
* 中断：进程在等待某个条件形成或者接受某个信号
   ps 工具表示的状态码： S    interruptible sleep (waiting for an event to complete)
* 不可中断：进程正处于内核态关键流程种，并且这些流程是不可打断。最常见的就是等待硬件设备的IO响应。例如，当一个进程向磁盘写入数据的时候，为了保证数据的一致性，
    在得到磁盘恢复之前，该进程是不能被其他进程中断或者打断.由此可见不可中断实际上是内核为进程和硬件设备的一种保护机制。
    ps 工具表示的转台码：D    uninterruptible sleep (usually IO)
* 僵死：进程已终止，但是进程描述符仍然存在，需要其父进程调用wait4()系统调用释放。
    ps 工具表示的状态码： Z a defunct("zombie") process
* 停止：进程收到控制信号(例如:SIGSTOP,SIGSTP,SIGTIN,SIGTOU等)后停止运行
    ps 工具表示的状态吗：T stopped by job control signal  

其他状态还包括W(),t(stopped by debugger during the tracing),X(dead (should never be seen)), 同时在BSD类型的操作系统下还包括以下集中特殊的状态：
    <    high-priority (not nice to other users)
    N    low-priority (nice to other users)
    L    has pages locked into memory (for real-time and custom IO)
    s    is a session leader
    l    is multi-threaded (using CLONE_THREAD, like NPTL pthreads do)
    +    is in the foreground process group
```

> 平均负责多少为合理?
> 对于这个问题其实并没有一个确定的答案。最理想的情况下，我们希望机器有多少个cpu core，同一时间就有多少个进程占用cpu使用。但实际生产中，我们发现当系统的平均负载超过了cpu core，服务的性能也没有受到影响。不过根据经验来看，当系统**平均负载高于cpu数量的70%**的时候，一般机器就会出现进程响应变慢，影响服务质量。在实际使用中最推荐的做法是**通过监控，观察平均负载的历史数据，判断负载的变化趋势来做分析更准确**。


> 平均负载与cpu利用率的关系：
> * cpu密集型服务： 平均负载越高，cpu利用率越高。
> * IO密集型服务：平均负载越高，cpu利用率可能很低。因为IO密集型服务在IO操作期间处于不可中断状态，此时进程不做运算操作，只是在等待硬件返回结果。如果IO响应很慢（使用性能低的磁盘），这时候我们看到系统的平均负载很高，但是cpu利用率不高。


> 查看系统cpu core数量
```
cat /proc/cpuinfo|grep 'model name' |wc -l
```

> stress 使用
```
# stress 是一个 Linux 系统压力测试工具,使用yum install -y stress 安装
# 模拟一个 CPU 使用率 100% 的场景
stress --cpu 56 --timeout 600
#模拟 I/O 压力，即不停地执行 sync
stress -i 56 --timeout 600
```

> sysstat 使用
```
sysstat 包含了常用的 Linux 性能工具，用来监控和分析系统的性能。使用yum install -y sysstat 安装
* mpstat 是一个常用的多核 CPU 性能分析工具，用来实时查看每个 CPU 的性能指标，以及所有 CPU 的平均指标。
* pidstat 是一个常用的进程性能分析工具，用来实时查看进程的 CPU、内存、I/O 以及上下文切换等性能指标。
# -P ALL 表示监控所有CPU，后面数字5表示间隔5秒后输出一组数据
mpstat -P ALL 5
# 间隔5秒后输出一组数据
pidstat -u 5 1
```

## CPU 上下文切换
