package basic

import (
	"fmt"
	"testing"
)

// 可以在 builtin.go 中看到下面的定义
// type byte = uint8
// type rune = int32
func TestType(t *testing.T) {
	s := "1" // string 和 rune 切片之间可以直接互相转换
	c := '1'
	fmt.Printf("s：%T\n", s)
	fmt.Printf("c：%T\n", c)

	// rune(int32) 可以直接转 byte(uint8)
	println(byte(c))

	fmt.Printf("en: %T, cn: %T\n", 'a', '啊')

	// 整型可以直接转浮点型
	println(float64(1))
}
