package basic

import (
	"testing"
)

/*
https://gfw.go101.org/article/nil.html

- nil 不是关键字，是一个 预声明标识符（变量名可以是 nil）
- 预声明标识符 nil 没有默认类型

- 不同类型的 nil 值的尺寸很可能不相同（unsafe.Sizeof）
- 不同类型的 nil 值不能比较（interface{} 类型比较特殊）

初学者在学习类型零值的时候，可以了解到 指针类型、接口类型 的零值是 nil；但是没有进一步了解这之间的联系和区别就会遇到各式各样的问题
*/

func TestNil(t *testing.T) {
	notKeywordTest(t)
}

func notKeywordTest(t *testing.T) {
	nil := 1
	t.Log(nil)
}
