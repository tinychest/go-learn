package cache_f

import "sync"

var allCache sync.Map

func set(name string, cache interface{}) {
	allCache.Store(name, cache)
}

func Get(name string) (interface{}, bool) {
	return allCache.Load(name)
}

// Refresh 如果后续有希望基于 HTTP 请求刷新缓存的需求，可以调用当前方法
func Refresh(name string) bool {
	v, ok := Get(name)
	if !ok {
		return false
	}
	if v == nil {
		return false
	}
	f, ok := v.(interface{ Refresh() })
	if ok {
		return false
	}
	f.Refresh()
	return true
}
