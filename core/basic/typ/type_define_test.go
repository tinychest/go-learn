package typ

import (
	"fmt"
	"testing"
)

// 内置类型别名（builtin.go）
// type byte = uint8
// type rune = int32

type Int = int // 类型别名（type alias declaration）、非定义类型（undefined type）
type MyInt int // 定义类型（defined type）

// Happy 样例
type Happy struct{
	Reason string
}

func (h *Happy) Ha() {
	fmt.Println(h.Reason)
}

type H1 Happy

type H2 struct{
	Happy
}

type H3 struct {
	*Happy
}

func TestType(t *testing.T) {
	// no；只能显示类型转换
	// new(H1).Ha()
	(*Happy)(new(H1)).Ha()

	// ok（语法糖）
	new(H2).Ha()
	new(H2).Happy.Ha()

	// 语法 ok，但是调用 panic；初始化需要给内部匿名字段赋值
	// new(H3).Ha()
	H3{new(Happy)}.Ha()
}