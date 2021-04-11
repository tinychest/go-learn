package lock

import (
	"sync"
	"time"
)

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.count++
	time.Sleep(cost)
}

func (l *Lock) Read() {
	l.mu.Lock()
	defer l.mu.Unlock()

	time.Sleep(cost)
	_ = l.count
}
