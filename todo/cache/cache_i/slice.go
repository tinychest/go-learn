package cache_i

import (
	"fmt"
	"sync"
	"time"
)

type SliceCacheable[T interface{}] interface {
	Query() ([]T, error)
}

// SliceWrapper Go 中匿名接口使用
type SliceWrapper[T interface{}] struct {
	Func func() ([]T, error)
}

func (w SliceWrapper[T]) Query() ([]T, error) {
	return w.Func()
}

type SliceCache[T interface{}] struct {
	once sync.Once
	mu   sync.RWMutex

	name       string
	dataSource SliceCacheable[T]
	data       []T
	lastUpdate int64
}

func NewSlice[T interface{}](name string, val SliceCacheable[T]) SliceCache[T] {
	c := SliceCache[T]{
		name:       name,
		dataSource: val,
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

	data, err := c.dataSource.Query()
	if err != nil {
		panic(fmt.Errorf("刷新缓存(%s)时出错：%w", c.name, err))
	}

	c.data = data
	c.lastUpdate = time.Now().Unix()
}
