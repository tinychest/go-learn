package _reflect

import (
	"reflect"
	"testing"
)

func TestTypeJudge(t *testing.T) {
	// ptrJudgeTest(t)
	// interfaceImplTest(t)
}

// 反射判断是否是指针类型
func ptrJudgeTest(t *testing.T) {
	var (
		ptr     *int
		unknown interface{} = ptr
	)

	v := reflect.ValueOf(unknown)
	t.Log(v.Kind() == reflect.Ptr)
}

// 反射判断是否实现了指定接口
func interfaceImplTest(t *testing.T) {
	type i interface { hello() }

	var s interface{}

	// 直接判断（不通过反射）
	_, ok := s.(i)
	t.Log(ok)

	// 这样写会引发 panic: reflect: nil type passed to Type.Implements
	// typ.Implements(reflect.TypeOf((i)(nil)))

	ok = reflect.TypeOf(s).Implements(reflect.TypeOf((*i)(nil)).Elem())
	t.Log(ok)
}
