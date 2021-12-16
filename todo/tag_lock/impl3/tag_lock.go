package impl3

import (
	"errors"
	"go-learn/todo/tag_lock"
	"sync"
)

type limiter sync.Map

func NewLimiter() tag_lock.StatusLocker {
	return new(limiter)
}

func (l *limiter) Lock(key interface{}) bool {
	_, ok := ((*sync.Map)(l)).LoadOrStore(key, struct{}{})
	return !ok
}

func (l *limiter) Unlock(key interface{}) error {
	_, ok := ((*sync.Map)(l)).LoadAndDelete(key)
	if !ok {
		return errors.New("unlock failed, nolock")
	}
	return nil
}
