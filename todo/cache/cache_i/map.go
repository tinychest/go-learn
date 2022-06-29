package cache_i

import (
	"fmt"
	"sync"
	"time"
)

type Key interface {
	string | int64
}

type MapCacheable[K Key, V interface{}] interface {
	Query() (map[K]V, error)
}

type MapCache[K Key, V interface{}] struct {
	once sync.Once
	mu   sync.RWMutex

	name       string
	dataSource MapCacheable[K, V]
	data       map[K]V
	lastUpdate int64
}

func (MapCache[K, V]) New(name string, val MapCacheable[K, V]) MapCache[K, V] {
	c := MapCache[K, V]{
		name:       name,
		dataSource: val,
	}
	_, ok := Get(name)
	if ok {
		panic(fmt.Errorf("duplication map cache name: %s", name))
	}
	set(name, c)
	return c
}

func (c *MapCache[K, V]) Load() map[K]V {
	c.once.Do(c.Refresh)

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data
}

func (c *MapCache[K, V]) Get(k K) V {
	return c.Load()[k]
}

func (c *MapCache[K, V]) Refresh() {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, err := c.dataSource.Query()
	if err != nil {
		panic(fmt.Errorf("刷新缓存(%s)时出错：%w", c.name, err))
	}

	c.data = data
	c.lastUpdate = time.Now().Unix()
}
