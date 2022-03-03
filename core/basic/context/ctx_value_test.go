package context

import (
	"context"
	"testing"
)

func TestValueContext(t *testing.T) {
	// oneTest(t)
	branchTest(t)
}

func oneTest(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "小明")
	ctx = context.WithValue(ctx, "userId", 1)

	userId, ok := ctx.Value("userId").(int)
	t.Log(userId, ok)

	name, ok := ctx.Value("name").(string)
	t.Log(name, ok)
}

func branchTest(t *testing.T) {
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, "key", "value1")
	ctx2 := context.WithValue(ctx, "key", "value2")

	t.Log(ctx1.Value("key"))
	t.Log(ctx2.Value("key"))
}
