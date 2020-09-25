package singleton

import "sync"

var lazyInstance *lazyCache
var once sync.Once

// 私有化lazyCache，限制用户可以直接实例化cache
type lazyCache struct {
	content sync.Map
}

// GetLazyCache 提供lazyCache的全局统一调用方式，由于使用了懒汉模式，
// 防止用户每次调用都会实例化lazyCache,使用once方法限制lazyCache整个启动期间这可以实例化一次
func GetLazyCache() *lazyCache {
	once.Do(func() {
		lazyInstance = &lazyCache{
			content: sync.Map{},
		}
	})
	return lazyInstance
}
