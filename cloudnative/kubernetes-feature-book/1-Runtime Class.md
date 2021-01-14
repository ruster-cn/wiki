# 1. Runtime Class

## Runtime Class 干啥的？

Runtime Class 出现的背景:
* 各种容器运行时的出现： runc,kata,gvisor 等。每种运行时都各有优势。runc更轻量，kata更安全。
* kubelet 开始支持CRI标准，对接多种容器运行时，不在单一的只能支持docker.

在这样的背景下，对于不通的业务，在部署的时候，想选择不同的运行时。为了支持业务部署以及运维方便，一个通俗的办法就是，业务部署时指定runtime,调度器将其调度到支持该runtime的节点。节点使用正确的runtime启动Pod。Runtime Class就是用来支持该功能的。

Runtime Class v1.14 开始支持默认打开，v1.16 支持单集群多种runtime class 共存。v1.18 支持Pod Overhead

## Runtime Class 如何使用？

* 确保kube-apiserver 和 kubelet的RuntimeClass feature gate 同时开启。
* 定义集群中存在的RuntimeClass，一个RuntimeClass定义对应于一个runtime。
* pod yaml在spce.runtimeClassName 指定需要使用的RuntimeClass。

![](image/1.drawio.svg)

如上所示：YAML 文件包含两个部分：上部分负责创建一个名字叫 runv 的 RuntimeClass 对象，下部分负责创建一个 Pod，该 Pod 通过 spec.runtimeClassName 引用了 runv 这个 RuntimeClass。



## Runtime Class 的原理
* kube-apiserver 接收到创建 Pod 的请求，根据runtime class定位为pod添加runtime的nodeselector；
* 方格部分表示三种类型的节点。每个节点上都有 Label 标识当前节点支持的容器运行时，节点内会有一个或多个 handler，每个 handler 对应一种容器运行时。比如第二个方格表示节点内有支持 runc 和 runv 两种容器运行时的 handler；第三个方格表示节点内有支持 runhcs 容器运行时的 handler；
* 调度器根据nodeSelector筛选节点完成调度: Pod 最终会调度到中间方格节点上，并最终由 runv handler 来创建 Pod。


## 其他知识
* pod Overhead(pod 额外开销):为什么要引入overhead呢。我们知道使用docker的时候，每个pod都有一个pause容器，这个pause容器功能简单基本没有资源消耗，因此在调度和kubelet准入的时候，都不计算这个容器的资源。但是对于kata这种安全容器来说，各种组件占用的资源加起来至少有100M，这就没办法忽略这部分资源了。除了在调度和kubelet准入的时候被用到之外，还会被计算在namespace 的quota和kubelet pod 驱逐中。


## 实操

### [kubelet 使用containerd CRI](https://github.com/containerd/containerd/blob/master/docs/installation.md)

* 安装 **seccomp** 所需的库: yum install libseccomp-devel

>seccomp 是啥？

* 在containerd 的[releases](https://github.com/containerd/containerd/releases) 下载对应版本的二进制

> 每种包的差别

* 解压安装包到根目录:  tar --no-overwrite-dir -C / -xzf cri-containerd-${VERSION}.linux-amd64.tar.gz
* 修改10-kubeadm.conf文件，为kubelet 添加如下启动参数: --container-runtime=remote --runtime-request-timeout=15m --container-runtime-endpoint=unix:///run/containerd/containerd.sock
* 启动containerd 和kubelet: systemctl restart containerd kubelet

## kataContainer 安装
https://github.com/kata-containers/documentation/blob/master/how-to/containerd-kata.md
https://github.com/kata-containers/documentation/blob/master/Developer-Guide.md#run-kata-containers-with-kubernetes



## build rootfs image


kata repo: http://download.opensuse.org/repositories/home:/katacontainers:/releases:/x86_64:/head/CentOS_7/

yum install kata-runtime

```
git clone https://github.com/kata-containers/osbuilder
$ export ROOTFS_DIR=${GOPATH}/src/github.com/kata-containers/osbuilder/rootfs-builder/rootfs
$ sudo rm -rf ${ROOTFS_DIR}
$ cd $GOPATH/src/github.com/kata-containers/osbuilder/rootfs-builder
$ script -fec 'sudo -E GOPATH=$GOPATH USE_DOCKER=true SECCOMP=no ./rootfs.sh ${distro}'
$ cd $GOPATH/src/github.com/kata-containers/osbuilder/image-builder
$ script -fec 'sudo -E USE_DOCKER=true ./image_builder.sh ${ROOTFS_DIR}'
```
rootfs vs initrd image 
https://www.sofastack.tech/blog/kata-container-2.0-road-to-attack/
https://mp.weixin.qq.com/s?__biz=MzUzOTk2OTQzOA==&mid=2247483874&idx=1&sn=cdc118f8c76a6bed6a6bd15153f5cb10&chksm=fac11313cdb69a055a2a200883b348a30f4d80f219b2f33a628efeccbfd6fd54efc7f7706f93&scene=21
https://cloud.tencent.com/developer/news/52377
https://blog.csdn.net/zhonglinzhang/article/details/86489695



参考
1. [从零开始入门 K8s：理解 RuntimeClass 与使用多容器运行时](https://www.infoq.cn/article/Ov2o7E3L1UbkCbq1V7o5)
2. [容器运行时类(Runtime Class)](https://kubernetes.io/zh/docs/concepts/containers/runtime-class/#cri-configuration)
