package _reflect

import . "reflect"

// 函数：指针解引用
func PtrUnReference(valuePtr *Value) {
	var value = *valuePtr

	for value.Kind() == Ptr {
		value = Indirect(value)
	}
	*valuePtr = value
}
