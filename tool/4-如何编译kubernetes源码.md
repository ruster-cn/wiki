# 如何编译kubernetes源码

- 二进制编译
```
make all
```
- 镜像编译
```
KUBE_BUILD_PLATFORMS=linux/amd64 KUBE_BUILD_CONFORMANCE=n KUBE_BUILD_HYPERKUBE=n make release-images
```
镜像发版产生的镜像在_output/release-images/amd64，为相关的tar文件，需要使用docker load -i 和docker tag 处理。
- 如何为kubernetes添加自己的依赖
  - 在staging src下加入需要引入的库
  - 在vendor下加入软连指向staging中加入的库
  - 修改go.mod,将依赖replace到staging中的库