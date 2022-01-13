package basic

import (
	"testing"
)

// 可以在 builtin.go 中看到下面的定义
// type byte = uint8
// type rune = int32
func TestType(t *testing.T) {
	s := "1" // string 和 rune 切片之间可以直接互相转换
	c := '1'
	t.Logf("S1：%T\n", s)
	t.Logf("c：%T\n", c)

	// rune(int32) 可以直接转 byte(uint8)
	t.Log(byte(c))

	t.Logf("en: %T, cn: %T\n", 'a', '啊')

	// 整型可以直接转浮点型
	t.Log(float64(1))
}
