package context

import (
	"context"
	"testing"
)

// 嵌套形式的数据添加
func TestValueContext(t *testing.T) {
	valueContext := context.WithValue(context.Background(), "name", "小明")

	valueContext = context.WithValue(valueContext, "userId", 1)

	process(valueContext)
}

func process(ctx context.Context) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		println("没有 userId")
	} else {
		println("取到 userId：", userId)
	}

	name, ok := ctx.Value("name").(string)
	if !ok {
		println("没有 name")
	} else {
		println("取到 name：", name)
	}
}
