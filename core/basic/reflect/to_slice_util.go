package _reflect

import (
	"fmt"
	. "reflect"
)

const (
	toString    = 0
	toInterface = 1
)

// 将参数中所有 基础字段 存储到一个 []string 作为返回值
func ToString(param interface{}) []string {
	return to(param, toString).([]string)
}

// 将参数中所有 基础字段 存储到一个 []interface 作为返回值
func ToInterface(param interface{}) []interface{} {
	return to(param, toInterface).([]interface{})
}

func to(param interface{}, model int) interface{} {
	// 最终结果
	strings := make([]string, 0)
	interfaces := make([]interface{}, 0)

	value, ok := param.(Value)
	if !ok {
		value = ValueOf(param)
	}
	PtrUnReference(&value)

	// 切片（结构体、字符串）、结构体、字符串
	switch value.Kind() {
	case Slice, Array:
		// 遍历切片、数组元素
		for index := 0; index < value.Len(); index++ {
			if model == toString {
				strings = append(strings, ToString(value.Index(index))...)
			}
			if model == toInterface {
				interfaces = append(interfaces, ToInterface(value.Index(index))...)
			}
		}
	case Struct:
		// 遍历结构体字段（没有指定忽略的）
		for index := 0; index < value.NumField(); index++ {
			field := value.Field(index)
			if value.Type().Field(index).Tag.Get("toString") != "-" {
				if model == toString {
					strings = append(strings, ToString(field)...)
				}
				if model == toInterface {
					interfaces = append(interfaces, ToInterface(field)...)
				}
			}
		}
	default:
		if model == toString {
			strings = append(strings, fmt.Sprint(value))
		}
		if model == toInterface {
			interfaces = append(interfaces, value.Interface())
		}
	}

	if model == toString {
		return strings
	}
	if model == toInterface {
		return interfaces
	}
	// never reach here
	return nil
}
