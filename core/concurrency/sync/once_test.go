package sync

import (
	"sync"
	"testing"
)

/*
sync.Once 在 Go 中是并发安全的只执行一次的逻辑实现
相关词条：DCL、懒汉式

Go DCL：https://turtledev.in/posts/go-concurrency-patterns-double-checked-locking/
*/

func TestOnce(t *testing.T) {
	_ = sync.Once{}
}
