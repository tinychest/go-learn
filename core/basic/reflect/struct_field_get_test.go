package _reflect

import (
	"fmt"
	"go-learn/core"
	"reflect"
	"testing"
	"unsafe"
)

func TestGetStructFieldValue(*testing.T) {
	p1 := core.Person{
		Age:  1,
		Name: "a",
	}
	p2 := core.Person{
		Age:  2,
		Name: "b",
	}

	// 获取 p1 的 Age
	v1 := reflect.ValueOf(p1)
	fmt.Println(v1.FieldByName("Age").Interface())
	// 获取 p2 的 Age
	v2 := reflect.ValueOf(p2)
	fmt.Println(v2.FieldByName("Age").Interface())

	// 你会发现，不像 java 能够利用一次 Class 的反射结果，对不同实例进行反射操作；Go 每次都要去获取一个 reflect.Value
	t := reflect.TypeOf(p1)
	ageField, _ := t.FieldByName("Age")
	ageFieldPtr := uintptr(unsafe.Pointer(&p1)) + ageField.Offset
	*((*int)(unsafe.Pointer(ageFieldPtr))) = 3
	fmt.Println(p1.Age)
	fmt.Println(*((*int)(unsafe.Pointer(ageFieldPtr))))
}
