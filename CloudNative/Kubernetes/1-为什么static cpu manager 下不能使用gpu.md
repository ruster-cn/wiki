# 1- 为什么static cpu manager下不能使用gpu

## 问题发现
当kubelet设置了--cpu-manager-policy=static参数后，使用了gpu的容器，在启动后不就会发现，服务在容器内失去了对gpu设备的使用权限。执行nvidia-smi将会报如下异常：
```
Failed to initialize NVML: Unknown Error
```
用trace跟踪一下，可以获取更明确的报错：
```
strace -v -a 100 -s 1000 nvidia-smi
close(3)                                                                                           = 0
open("/dev/nvidiactl", O_RDWR)                                                                     = -1 EPERM (Operation not permitted)
open("/dev/nvidiactl", O_RDONLY)                                                                   = -1 EPERM (Operation not permitted)
fstat(1, {st_dev=makedev(0, 704), st_ino=4, st_mode=S_IFCHR|0620, st_nlink=1, st_uid=0, st_gid=5, st_blksize=1024, st_blocks=0, st_rdev=makedev(136, 1), st_atime=2019/04/23-17:35:28.678347231, st_mtime=2019/04/23-17:35:28.678347231, st_ctime=2019/04/23-17:33:09.682347235}) = 0
write(1, "Failed to initialize NVML: Unknown Error\n", 41Failed to initialize NVML: Unknown Error
)                                         = 41
exit_group(255)                                                                                     = ?
+++ exited with 255 +++

```

当我们比较正常容器和异常容器的device.list 文件发现了不同，异常容器的device.list少了nvidia设备。

```
正常容器的devices.list文件内容
c 1:5 rwm
c 1:3 rwm
c 1:9 rwm
c 1:8 rwm
c 5:0 rwm
c 5:1 rwm
c *:* m
b *:* m
c 1:7 rwm
c 136:* rwm
c 5:2 rwm
c 10:200 rwm
c 195:255 rw
c 195:3 rw

异常容器的devices.list文件内容
c 1:5 rwm
c 1:3 rwm
c 1:9 rwm
c 1:8 rwm
c 5:0 rwm
c 5:1 rwm
c *:* m
b *:* m
c 1:7 rwm
c 136:* rwm
c 5:2 rwm
c 10:200 rwm
```
我们可以看下缺少的两项device,就是gpu卡和nvidia-smi 使用到的 nvidiactl
![](image/1.drawio.svg)

## 问题追踪
为什么会出现奇怪的问题呢，kubelet加了一个--cpu-manager-policy=static的参数，会导致容器运行过程中丢失设备。要搞清楚这个问题肯定得从两方面查一下，一个是增了--cpu-manager-policy=static参数后kubelet的工作流发生了那些变化，这些变化又是如何影响到底层容器的。

--cpu-manager-policy=static的功能是啥？默认情况下kubelet创建的pod都是通过CFS的方式来分配使用物理机的cpu资源。而static cpu manager提供了cpu set的功能。能够给某些container绑定指定的cpus，达到绑定cpus的目标，提升cpu敏感型任务的性能。按照线上生产环境的数据显示container如果使用了cpu set，业务的性能提升在15%-26%左右。

static cpu manager是如何工作的？

## 修复方法




https://cloud.tencent.com/developer/article/1402119
https://www.cnblogs.com/sparkdev/p/9129334.html