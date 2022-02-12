package di

import (
	"go.uber.org/dig"
	"testing"
)

type A struct {
	Value string
	b *B
}

func NewA(b *B) *A {
	return &A{ Value: "a", b: b }
}

type B struct{
	Value string
	c *C
}

func NewB(c *C) *B {
	return &B{ Value: "b", c: c }
}

type C struct{
	Value string
}

func NewC() *C {
	return &C{ Value: "c" }
}

// 常规获取方式（每次都要执行相同的操作）
func TestSimple(t *testing.T) {
	c := NewC()
	b := NewB(c)
	a := NewA(b)

	t.Log(a)
}

// 通过预定义的依赖关系，直接获取实例（只要定义好了关系，之后直接获取即可）
// - 多次调用，获取实例的地址是相同的
// - 当前上下文示例比较简单，实际支持多参
func TestDI(t *testing.T) {
	d := dig.New()

	if err := d.Provide(NewA); err != nil {
		t.Fatal(err)
	}
	if err := d.Provide(NewB); err != nil {
		t.Fatal(err)
	}
	if err := d.Provide(NewC); err != nil {
		t.Fatal(err)
	}

	err := d.Invoke(func(a *A) {
		t.Log("获取到 A", a)
	})
	if err != nil {
		t.Fatal("获取 A 失败", err)
	}

	err = d.Invoke(func(a *A) {
		t.Log("获取到 A", a)
	})
	if err != nil {
		t.Fatal("获取 A 失败", err)
	}
}