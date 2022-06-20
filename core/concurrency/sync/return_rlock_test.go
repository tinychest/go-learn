package sync

import (
	"sync"
	"testing"
)

type container struct {
	sync.RWMutex
	i    int
	list []int
}

func newContainer() *container {
	return &container{
		i:    0,
		list: make([]int, 5),
	}
}

func (c *container) Get() []int {
	c.RLock()
	defer c.RUnlock()

	return c.list
}

func (c *container) Set() {
	c.Lock()
	defer c.Unlock()

	c.i++
	c.list = []int{c.i, c.i + 1, c.i + 2, c.i + 3, c.i + 4}
}

// go test -v --race -run \QTestSingleTest\E
// Get 方法中的读锁如果注释掉，从现象上说，并没有数据逻辑上的问题（都是连续的），但是会提示 DATA RACE
func TestSingleTest(t *testing.T) {
	c := newContainer()

	go func() {
		for {
			c.Set()
		}
	}()

	for i := 0; i < 100; i++ {
		t.Log(c.Get())
	}
	c.Get()
}
