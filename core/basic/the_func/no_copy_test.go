package the_func

import (
	"sync"
	"testing"
)

// Go 为什么要设计成 Xxx 和 *Xxx 都是实现了 (x Xxx) 方法，而只有 *Xxx 实现了 (x *Xxx) 方法，而 Xxx 不算，因为有些方法的实现机理就是要求改变 Receiver 的某些属性值
// （如，sync.Mutex 的 Lock 和 Unlock 方法的 Receiver 类型都是 *，否则，值拷贝将导致锁一经方法传递进行复制，就不可能实现对加的锁进行解锁了）
// 利用好这个 vet check 规则
// 详见1 cond.noCopy
// 详见2 google.golang.org\protobuf@v1.27.1\internal\pragma\pragma.go
// 参见 https://golang.org/issues/8005
// 参见 https://go-review.googlesource.com/c/go/+/22015/

// 实现 1
type DoNotCopy [0]sync.Mutex
type s1 struct {
	DoNotCopy
}

// 实现 2
type NoCopy struct{}
func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}

type s2 struct {
	NoCopy
}

func TestNoCopy(t *testing.T) {
	a := s1{}
	b := s2{}

	// 灰色警告的意思是，复制示例的类型定义中包含 sync.Mutex，这是 sync.Locker 类型，不应该复制
	// sync.Cond、sync.Pool 都
	_ = a
	_ = b
}
