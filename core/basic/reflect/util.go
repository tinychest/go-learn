package _reflect

import . "reflect"

// PtrUnReference 指针解引用
func PtrUnReference(ptr *Value) {
	var v = *ptr

	for v.Kind() == Ptr {
		v = Indirect(v)
	}
	*ptr = v
}