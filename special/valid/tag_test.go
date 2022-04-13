package valid

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

type Out struct {
	// 不加标签 会校验 子结构体内部带有标签的字段（比较特殊）
	// 加 validate:"required" 标签 会校验
	// 加 validate:""         标签 会校验
	// 加 validate:"-"        标签 不会校验
	In In `validate:"required"`
	// 加 validate:"required" 标签 对于 bool 类型的来说，值必须是 true 才能通过校验（因为 false 是零值）
	// 不加标签不会校验
	OK bool
	// 不会对结构体类型进行校验，只能够对其对应的指针类型进行校验
	Bot struct{ Name string } `validate:"required"`
}

type In struct {
	Name string `validate:"required"`
	// 不加标签不会校验
	OK bool
}

func TestStructTag(t *testing.T) {
	var o = Out{OK: true, In: In{Name:"1"}}
	err := validator.New().Struct(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
