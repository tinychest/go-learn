package util

import (
	"fmt"
	. "reflect"
	"strings"
)

// 打印任意类型的 切片 的详细信息
// 说明：这个完全是为学习阶段的 debug 方便，生产中不可能使用，一个反射的性能额外消耗，一个消耗空间，一个遍历的性能低下
func PrintSliceInfo(param interface{}) {
	printSliceInfo(sliceToInterfaceSlice(param), param)
}

func sliceToInterfaceSlice(param interface{}) []interface{} {
	// 一、参数处理
	// 1、反射值处理
	value, ok := param.(Value)
	if !ok {
		value = ValueOf(param)
	}
	// 2、指针引用处理
	PtrUnReference(&value)

	// 二、参数校验
	if value.Kind() != Slice {
		// 参照 Go 源码的类型错误异常
		panic(&ValueError{Method: "PrintSlice", Kind: value.Kind()})
	}

	slice := make([]interface{}, 0, value.Len())
	for index := 0; index < value.Len(); index++ {
		slice = append(slice, value.Index(index).Interface())
	}

	return slice
}

// 制定打印详细格式的打印方法
// param：打印原参数的地址
func printSliceInfo(slice []interface{}, param interface{}) {
	fmt.Printf("[%d/%d] %p %-11s\n", len(slice), cap(slice), param, sprintSlice(slice))
}

// 返回切片指定格式字符串的打印内容
func sprintSlice(slice []interface{}) string {
	builder := strings.Builder{}

	// 参照 print.go/doPrintln
	builder.WriteString("[")
	for index := 0; index < len(slice); index++ {
		if index > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprint(slice[index]))
	}
	builder.WriteString("]")

	return builder.String()
}
