package _reflect

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestSizeof(t *testing.T) {
	// int 8
	t.Log(unsafe.Sizeof(reflect.Int))
	// rune = Unint8 8
	t.Log(unsafe.Sizeof(reflect.Uint8))
	// string 虽然支持该类型，但是并得不到具体的类型大小
	t.Log(unsafe.Sizeof("012345678901234567890123123123123123123"))
	// slice（无论 len cap 是多少，这里返回的都是常量）
	t.Log(unsafe.Sizeof(make([]string, 2, 4)))
}
