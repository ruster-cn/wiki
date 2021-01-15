# workqueue in k8s client-go 

## work queue

接口定义
```Golang
type Interface interface {
	Add(item interface{})
	Len() int
	Get() (item interface{}, shutdown bool)
	Done(item interface{})
	ShutDown()
	ShuttingDown() bool
}
```
