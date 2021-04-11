package _reflect

import (
	"go-learn/core"
	. "reflect"
	"testing"
	"unsafe"
)

// 利用反射修改结构体实例字段的值
func TestReflectGetField(t *testing.T) {
	person1 := &core.Person{Age: 10, Name: "小明"}
	person2 := &core.Person{Age: 11, Name: "小红"}

	// 实际情况可能是接口类型为 interface{}，当确定为指定的接口体类型，就做特定的操作

	nameField, _ := TypeOf(person1).Elem().FieldByName("Name")

	namePtr1 := uintptr(unsafe.Pointer(person1)) + nameField.Offset
	namePtr2 := uintptr(unsafe.Pointer(person2)) + nameField.Offset

	*((*string)(unsafe.Pointer(namePtr1))) = "大明"
	*((*string)(unsafe.Pointer(namePtr2))) = "大红"

	println(person1)
	println(person2)
}
