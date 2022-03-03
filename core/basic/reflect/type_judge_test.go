package _reflect

import (
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
	// ptrIsTest(t)
	interfaceImplTest(t)
	// interfaceIsTest(t)
}

func ptrIsTest(t *testing.T) {
	var (
		ptr     *int
		unknown interface{} = ptr
	)

	v := reflect.ValueOf(unknown)
	t.Log(v.Kind() == reflect.Ptr)
}

func interfaceImplTest(t *testing.T) {
	type i interface {
		hello()
	}

	typ := reflect.TypeOf("")
	// 这样写会引发 panic: reflect: nil type passed to Type.Implements
	// typ.Implements(reflect.TypeOf((i)(nil)))
	t.Log(typ.Implements(reflect.TypeOf((*i)(nil)).Elem()))
}

func interfaceIsTest(t *testing.T) {
	type i interface {
		hello()
	}
	var s interface{}

	_, ok := s.(i)
	t.Log(ok)
}
