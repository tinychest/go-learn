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
	var sumPtr *int
	var unknown interface{} = sumPtr

	val := reflect.ValueOf(unknown)

	if val.Kind() == reflect.Ptr {
		println("是指针类型")
	} else {
		println("不是指针类型")
	}
}
