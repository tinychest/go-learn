package typ

import (
	"testing"
	"unsafe"
)

// Struct → Interface ✔
// []Struct → []Interface ❌
// []Struct 并不能简单通过 unsafe 操作得到对应 []Interface，只能老实创建一个接口切片，循环遍历添加
// TODO 需要给出一个能信服的原因

type Killer interface {
	Kill()
}

type Assassin struct{}

func (a Assassin) Kill() {}

func TestUnsafeSliceCast(t *testing.T) {
	as := []Assassin{{}}
	rollCall(*(*[]Killer)(unsafe.Pointer(&as)))
}

func rollCall(ks []Killer) {
	for _, k := range ks {
		// panic
		k.Kill()
	}
}
