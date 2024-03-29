package todo

import (
	"github.com/jinzhu/copier"
	"go-learn/tool"
	"testing"
)

// TODO 可以从这个库学习反射的语法、基础知识
func TestCopier(t *testing.T) {
	s1 := []string{"1", "2", "3"}
	var s2 []int

	// 非指针 → err
	// 类型不匹配 → no error、not copy
	if err := copier.Copy(&s2, s1); err != nil {
		t.Fatal(err)
	}
	tool.PrintSlice(s1)
	tool.PrintSlice(s2)
}
