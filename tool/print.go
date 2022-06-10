package tool

import (
	"encoding/json"
	"fmt"
	. "reflect"
	"strings"
)

// PrintMap 打印 map
// 其实使用 json 序列化的方式去打印，最简单，但是 json 只支持 string 类型的键
func PrintMap(m interface{}) {
	raw := fmt.Sprint(m)[3:]
	raw = strings.ReplaceAll(raw, "[", "{")
	raw = strings.ReplaceAll(raw, "]", "}")
	fmt.Println(raw)
}

// PrintAddr 打印地址
func PrintAddr(p interface{}) {
	fmt.Printf("%p\n", p)
}

// PrintSlice 打印任意类型的 切片 的详细信息
// 说明：这个完全是为学习阶段的 debug 方便，生产中不可能使用，一个反射的性能额外消耗，一个消耗空间，一个遍历的性能低下
func PrintSlice(param interface{}) {
	printSlice(sliceToInterfaceSlice(param), param)
}

// PrintJSON 序列化成 json 打印
func PrintJSON(p interface{}) {
	fmt.Println(MustMarshalJSON(p))
}

func MustMarshalJSON(p interface{}) []byte {
	bs, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return bs
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

	slice := make([]interface{}, 0, value.Cap())
	for i := 0; i < value.Len(); i++ {
		slice = append(slice, value.Index(i).Interface())
	}

	return slice
}

// 制定打印详细格式的打印方法
// param：打印原参数的地址
func printSlice(slice []interface{}, param interface{}) {
	fmt.Printf("[%d/%d] %p %-11s\n", len(slice), cap(slice), param, sprintSlice(slice))
}

// 返回切片指定格式字符串的打印内容
func sprintSlice(slice []interface{}) string {
	b := strings.Builder{}

	// 参照 print.go/doPrintln
	b.WriteString("[")
	for index := 0; index < len(slice); index++ {
		if index > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprint(slice[index]))
	}
	b.WriteString("]")

	return b.String()
}
