package typ

import (
	"sync"
	"testing"
)

// Go 为什么要设计成 Xxx 和 *Xxx 都是实现了 (x Xxx) 方法，而只有 *Xxx 实现了 (x *Xxx) 方法，而 Xxx 不算，因为有些方法的实现机理就是要求改变 Receiver 的某些属性值
// （如，sync.Mutex 的 Lock 和 Unlock 方法的 Receiver 类型都是 *，否则，值拷贝将导致锁一经方法传递进行复制，就不可能实现对加的锁进行解锁了）
//
// 假如有个结构体，无论是因为业务逻辑，还是因为数据量过大的原因，你都不希望对其有复制的操作
// 那么你可以利用好 vet check 的 nocopy 规则
//
// 详见1 cond.noCopy
// 详见2 google.golang.org\protobuf@v1.27.1\internal\pragma\pragma.go
// 参见 https://golang.org/issues/8005
// 参见 https://go-review.googlesource.com/c/go/+/22015/

// 实现 1
type NoCopy1 [0]sync.Mutex

// 实现 2
type NoCopy2 struct{}
func (*NoCopy2) Lock()   {}
func (*NoCopy2) Unlock() {}

type s1 struct {
	NoCopy1
}

type s2 struct {
	NoCopy2
}

func TestNoCopy(t *testing.T) {
	a := s1{}
	b := s2{}

	// 灰色警告的意思是，复制示例的类型定义中包含 sync.Mutex，这是 sync.Locker 类型，不应该复制
	// sync.Cond、sync.Pool 都
	copyA := a
	copyB := b

	_ = copyA
	_ = copyB
}
