package basic

import (
	"testing"
)

// Go 中的 string 也是不可变的，同 Java
// 所以做字符串拼接的性能表现上：
//  预分配内存(cap)的[]byte > []byte ~> strings.Builder（推荐） ~> bytes.Buffer >> 直接 string 拼接
// 1.不要使用 + 的拼串 或 fmt.Sprintf
// 2.推荐使用预分配内存（strings.Builder.Grow(int)）的 strings.Builder 性能最好（比预分配内存的 []byte 少了一次 []byte 到 string 的转化）

// 说到 string 就离不开字符，string 可以理解为字符切片（因不可变性，理解为字符数组可能更恰当一些）
func TestString(t *testing.T) {
	// 1、支持同 切片 数组的范围取值转化
	// var str = "123"
	// str = str[:len(str) - 1]

	// 2、string 通过下标取值的类型是：uint8(byte)（所以 string 可以转化成：[]byte）
	// var value = str[0]

	// 3、string 通过 for range 语法遍历值类型是：int32(rune)
	// for _, value := range str {}

	// 4、string 是可以比较的（按照字典顺序）
}
