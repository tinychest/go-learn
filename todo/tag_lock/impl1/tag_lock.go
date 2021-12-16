package impl1

import (
	"errors"
	"go-learn/todo/tag_lock"
	"sync"
)

type locker sync.Map

func NewLocker() tag_lock.SyncLocker {
	return new(locker)
}

func (l *locker) Lock(key interface{}) {
	c, _ := ((*sync.Map)(l)).LoadOrStore(key, make(chan struct{}, 1))
	c.(chan struct{}) <- struct{}{}
}

func (l *locker) Unlock(key interface{}) error {
	c, ok := ((*sync.Map)(l)).Load(key)
	if !ok {
		return errors.New("unlock failed, nolock")
	}
	<-c.(chan struct{})
	return nil
}