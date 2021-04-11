package context

import (
	"context"
	"testing"
)

// context.WithValue 方法就是进行多个键值对数据的链接的
// context.Value     方法就是
func TestValueContext(t *testing.T) {
	valueContext := context.WithValue(context.Background(), "name", "小明")

	valueContext = context.WithValue(valueContext, "userId", 1)

	process(valueContext)
}

func process(ctx context.Context) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		println("没有 userId")
		return
	} else {
		println("取到 userId：", userId)
	}

	name, ok := ctx.Value("name").(string)
	if !ok {
		println("没有 name")
		return
	} else {
		println("取到 name：", name)
	}
}
