package cache_i

import (
	"fmt"
	"sync"
	"time"
)

type Cacheable[T interface{}] interface {
	Query() (T, error)
}

type Cache[T interface{}] struct {
	once sync.Once
	mu   sync.RWMutex

	name       string
	dataSource Cacheable[T]
	data       T
	lastUpdate int64
}

func New[T interface{}](name string, dataSource Cacheable[T]) Cache[T] {
	c := Cache[T]{
		name:       name,
		dataSource: dataSource,
	}
	_, ok := Get(name)
	if ok {
		panic(fmt.Errorf("duplication cache name: %s", name))
	}
	set(name, c)
	return c
}

func (c *Cache[T]) Load() T {
	c.once.Do(c.Refresh)

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data
}

func (c *Cache[T]) Refresh() {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, err := c.dataSource.Query()
	if err != nil {
		panic(fmt.Errorf("刷新缓存(%s)时出错：%w", c.name, err))
	}

	c.data = data
	c.lastUpdate = time.Now().Unix()
}
