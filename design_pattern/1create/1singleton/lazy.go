package singleton

import "sync"

/* 懒汉式 DCL */
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = new(Singleton)
		})
	}
	return lazySingleton
}
