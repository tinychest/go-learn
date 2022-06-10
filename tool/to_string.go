package tool

import . "reflect"

// ToString 将参数中所有 string 类型的字段值存储到一个 []string 作为返回值
// 参数类型支持：字符串、结构体[R]、切片（结构体[R]、字符串，不要求切片中的元素类型是一致的） - R：Recursive
// 宏观来说，目前只支持 string 基础类型
func ToString(param interface{}) []string {
	// 定义空返回值，不返回 nil
	var emptyResult []string

	// 最终结果
	strings := make([]string, 0)
	value, ok := param.(Value)
	if !ok {
		value = ValueOf(param)
	}
	PtrUnReference(&value)

	// 切片（结构体、字符串）、结构体、字符串
	switch value.Kind() {
	case Array, Slice:
		// 遍历数组元素
		for index := 0; index < value.Len(); index++ {
			strings = append(strings, ToString(value.Index(index))...)
		}
	case Struct:
		// 遍历结构体字段
		for index := 0; index < value.NumField(); index++ {
			strings = append(strings, ToString(value.Field(index))...)
		}
	case String:
		strings = append(strings, value.String())
	default:
		return emptyResult
	}

	return strings
}
