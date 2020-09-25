package singleton

import "sync"

var instance = &cache{
	content: sync.Map{},
}

// 限制用户可以直接实例化cache
type cache struct {
	content sync.Map
}

// GetCache 提供单例全局唯一的访问方法
func GetCache() *cache {
	return instance
}

func (c *cache) Add(key string, value interface{}) {
	c.content.Store(key, value)
}

func (c *cache) Get(key string) interface{} {
	v, ok := c.content.Load(key)
	if !ok {
		return nil
	}
	return v
}
