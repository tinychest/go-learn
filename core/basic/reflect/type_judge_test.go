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
	// ptrTest()
	interfaceTest()
	interfaceRecommendTest()
}

func ptrTest() {
	var sumPtr *int
	var unknown interface{} = sumPtr

	val := reflect.ValueOf(unknown)

	ok := val.Kind() == reflect.Ptr
	fmt.Println(ok)
}

func interfaceTest() {
	type i interface {
		hello()
	}
	var s interface{}

	var ok bool
	if s == nil {
		ok = false
	} else {
		sType := reflect.TypeOf(s)
		ok = sType.Implements(reflect.TypeOf((*i)(nil)).Elem())
	}
	fmt.Println(ok)
}

func interfaceRecommendTest() {
	type i interface {
		hello()
	}
	var s interface{}

	_, ok := s.(i)
	fmt.Println(ok)
}

