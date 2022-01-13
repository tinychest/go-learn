package _reflect

import (
	"go-learn/core"
	"reflect"
	"testing"
	"unsafe"
)

func TestGetStructFieldValue(t *testing.T) {
	// reflectGetTest(t)
	unsafeGetTest(t)
}

func reflectGetTest(t *testing.T) {
	p1 := core.Person{Age: 1, Name: "a"}
	p2 := core.Person{Age: 2, Name: "b"}

	v1 := reflect.ValueOf(p1)
	v2 := reflect.ValueOf(p2)

	v1AgeField := v1.FieldByName("Age")
	v1NameField := v1.FieldByName("Name")
	v2AgeField := v2.FieldByName("Age")
	v2NameField := v2.FieldByName("Name")

	t.Log(v1AgeField.Int())
	t.Log(v2AgeField.Interface())
	t.Log(v1NameField.String())
	t.Log(v2NameField.Interface())
}

func unsafeGetTest(t *testing.T) {
	p := core.Person{Age:  1, Name: "a"}

	// 你会发现，不像 java 能够利用一次 Class 的反射结果，对不同实例进行反射操作；Go 每次都要去获取一个 reflect.Value
	// 你会发现，Offset 是 StructField 类型的字段，StructField 类型是 reflect.Type 类型的方法（结构体字段的 Offset 只和类型相关）
	//   string 并不是一个基础类型，底层是 字节数组、其他长度 字段，数组指地址，因此 string 可以做到定长

	// 方式一、借助反射获取到的 offset
	nameField, _ := reflect.TypeOf(p).FieldByName("Name")
	nameAddr := uintptr(unsafe.Pointer(&p)) + nameField.Offset
	t.Log(*(*string)(unsafe.Pointer(nameAddr)))

	// 方式二、使用 unsafe
	nameAddr = uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.Name)
	t.Log(*(*string)(unsafe.Pointer(nameAddr)))

	// 方式三、使用 unsafe.Add 方法
	t.Log(*(*string)(unsafe.Add(unsafe.Pointer(&p), unsafe.Offsetof(p.Name))))
}
