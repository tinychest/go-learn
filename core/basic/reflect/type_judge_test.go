package _reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// 基础类型就不说了，还支持如下类型：
// Array
// Chan
// Func
// Interface
// Map
// Ptr
// Slice
// Struct
// UnsafePointer
func TestTypeJudge(t *testing.T) {
	ptrIsTest()
	interfaceImplTest()
	interfaceIsTest()
}

func ptrIsTest() {
	var (
		ptr     *int
		unknown interface{} = ptr
	)

	v := reflect.ValueOf(unknown)
	fmt.Println(v.Kind() == reflect.Ptr)
}

func interfaceImplTest() {
	type i interface {
		hello()
	}
	var unknown interface{}

	if unknown == nil {
		fmt.Println(false)
		return
	}

	t := reflect.TypeOf(unknown)
	ok := t.Implements(reflect.TypeOf((*i)(nil)).Elem())
	fmt.Println(ok)
}

func interfaceIsTest() {
	type i interface {
		hello()
	}
	var s interface{}

	_, ok := s.(i)
	fmt.Println(ok)
}
