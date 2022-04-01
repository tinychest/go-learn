package valid

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

type Out struct {
	// 不加 validate 标签会校验 In.Name
	// 加 validate:"required" 标签会校验 In.Name
	// 加 validate:"" 标签会校验 In.Name
	// 加 validate:"-" 标签不会校验 In.Name
	In In `validate:"-"`
}

type In struct {
	Name string `validate:"required"`
}

func TestStructTag(t *testing.T) {
	var o Out
	err := validator.New().Struct(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}