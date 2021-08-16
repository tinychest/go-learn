package lock

import (
	"sync"
	"time"
)

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.count++
	time.Sleep(cost)
}

func (l *RWLock) Read() {
	l.mu.RLock()
	defer l.mu.RUnlock()

	_ = l.count
	time.Sleep(cost)
}
