package _reflect

import (
	"go-learn/core"
	. "reflect"
	"testing"
	"unsafe"
)

func TestReflectSetField(t *testing.T) {
	// reflectSetTest(t)
	unsafeSetTest(t)
}

func reflectSetTest(t *testing.T) {
	p := &core.Person{Age: 10, Name: "小明"}

	ValueOf(p).Elem().FieldByName("Name").Set(ValueOf("大明"))

	t.Log(p)
}

func unsafeSetTest(t *testing.T) {
	p := &core.Person{Age: 10, Name: "小明"}

	// 方式一、借助反射获取到的 offset
	nameField, _ := TypeOf(p).Elem().FieldByName("Name")
	nameAddr := uintptr(unsafe.Pointer(&p)) + nameField.Offset
	*(*string)(unsafe.Pointer(nameAddr)) = "大明"
	t.Log(p)

	// 方式二、使用 unsafe
	nameAddr = uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.Name)
	*(*string)(unsafe.Pointer(nameAddr)) = "大明"
	t.Log(p)

	// 方式三、使用 unsafe.Add 方法
	*(*string)(unsafe.Add(unsafe.Pointer(&p), unsafe.Offsetof(p.Name))) = "大明"
	t.Log(p)
}