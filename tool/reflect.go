package tool

import . "reflect"

// 示例，仅用于演示常见的处理类型 以及 说明不同长度的基础类型 Go 都会以最大的来处理
// func ActualValueExample(value Value) interface{} {
//     switch value.Kind() {
//     case Bool:
//         return value.Bool()
//     case Int,Int8,Int16,Int32,Int64:
//         return value.Int()
//     case Uint,Uint8,Uint16,Uint32,Uint64:
//         return value.Uint()
//     case Float32,Float64:
//         return value.Float()
//     case String:
//         return value.String()
//     case Interface:
//         return value.Interface()
//     case Array, Slice:
//         return value.Slice(0, value.Len())
//     default:
//         panic("别以为什么牛鬼蛇神的类型我都支持")
//     }
// }

// PtrUnReference 指针解引用
func PtrUnReference(valuePtr *Value) {
	if valuePtr == nil {
		panic("util.PtrUnReference: can not deal with the nil Ptr")
	}

	var value = *valuePtr

	for value.Kind() == Ptr {
		value = Indirect(value)
	}
	*valuePtr = value
}
