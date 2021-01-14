# kubebuilder 实战

## 安装
1. 从[kubebuild release page](https://github.com/kubernetes-sigs/kubebuilder/releases)下载kubebuilder,解压后，将bin目录中文件拷贝到PATH目录中。
2. 执行 `curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash` 下载 kustomize,并拷贝到PATH目录中。
3. 下载[controller-gen](https://github.com/kubernetes-sigs/controller-tools/releases) 编译后，拷贝到PATH目录中。
## example 1: cronjob

1. 创建项目 

```shell
#创建项目目录
mkdir test-cronjob
cd  test-cronjob
#初始化gomod
go mod init tutorial.com/cronjob
#创建kubebuilder项目
kubebuilder init --domain tutorial.kubebuilder.io
```

2. 添加API

```shell
# 添加一个resource 
kubebuilder create api --group batch --version v1 --kind CronJob
```
执行以上命令，kubebuilder 会在项目跟目录生成api和controllers两个目录，分别存储api定义和controller。

Resource API定义的一些原则：
>
> 1. spec代表期望的状态，用户的任何输入都应该定义到spec中
> 2. status是实际的运行状态
> 3. API定义中的所有字段json后必须是驼峰格式
> 4. 字段理论上可以是任何基本类型，但是我们对数字类型有些特殊的约束：整数类型必须使用过int32 和 int64 类型。小数使用resource.Quantity 类型。
> 5. 基本类型一般都是有默认值的，所以对于一些需要强制用户填写的自动，应该使用基本类型的指针类型。

3. 编写controller

controller的工作就是协调给定的对象，使对象的真实状态与用户定义的spec中期望的状态一致。通常每个控制钱应该只专注于一个根kind，但也可能与其他的Kind交互。

kubebuilder会帮助生成大部分的结构，我们只需要更具需要补充CronJobReconciler struct和实现Reconcile方法即可。当Reconcile方法返回一个空结果，并且没有err的时候，说明本次协调成功。

4. 添加webhook 

我们还可以为资源实现一个 admission webhooks，在用户提交资源使，为其填充默认值和做一些值的校验工作。

```
# --defaulting: 默认值填充功能
# --programmatic-validation: 值校验功能
kubebuilder create webhook --group batch --version v1 --kind CronJob --defaulting --programmatic-validation --conversion

```


## 扩展

1. 

参考资料:
1. [kubebuilder introduction 中文版](https://cloudnative.to/kubebuilder/introduction.html)
2. [kubebuilder introduction 原版](https://book.kubebuilder.io/quick-start.html)