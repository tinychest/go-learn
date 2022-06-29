package cache_f

import (
	"fmt"
	"sync"
	"time"
)

type SliceDataFunc[T interface{}] func() ([]T, error)

type SliceCache[T interface{}] struct {
	once sync.Once
	mu   sync.RWMutex

	name       string
	dataFunc   SliceDataFunc[T]
	data       []T
	lastUpdate int64
}

func New[T interface{}](name string, val SliceDataFunc[T]) SliceCache[T] {
	c := SliceCache[T]{
		name:     name,
		dataFunc: val,
	}
	_, ok := Get(name)
	if ok {
		panic(fmt.Errorf("duplication slice cache name: %s", name))
	}
	set(name, c)
	return c
}

func (c *SliceCache[T]) Load() []T {
	c.once.Do(c.Refresh)

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data
}

func (c *SliceCache[T]) Filter(f func(T) bool) []T {
	list := c.Load()

	if f == nil {
		panic(fmt.Errorf("%s SliceCache filter, filter func can not be nil", c.name))
	}

	var res = make([]T, 0, len(list))
	for _, v := range list {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func (c *SliceCache[T]) Refresh() {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, err := c.dataFunc()
	if err != nil {
		panic(fmt.Errorf("刷新缓存(%s)时出错：%w", c.name, err))
	}

	c.data = data
	c.lastUpdate = time.Now().Unix()
}
