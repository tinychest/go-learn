package context

import (
	"context"
	"testing"
)

// 嵌套形式的数据添加
func TestValueContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "小明")
	ctx = context.WithValue(ctx, "userId", 1)

	process(t, ctx)
}

func process(t *testing.T, ctx context.Context) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		t.Log("没有 userId")
	} else {
		t.Log("取到 userId：", userId)
	}

	name, ok := ctx.Value("name").(string)
	if !ok {
		t.Log("没有 name")
	} else {
		t.Log("取到 name：", name)
	}
}
