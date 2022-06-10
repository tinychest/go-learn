package gomonkey

import (
	"github.com/agiledragon/gomonkey/v2"
	"testing"
)

func Hello() string {
	return "hello"
}

func TestSimpleFunc(t *testing.T) {
	specialHandle()

	res := Hello()
	t.Log(res)
}

// NOTE 直接 RUN 打桩不生效，只有以 DEBUG 模式运行才生效
// 参见官方文档：https://github.com/agiledragon/gomonkey
//
// 通过 GitHub 官方文档了解到，应该禁用 inline 才能让打桩生效
// 运行时，通过添加 -gcflags=all=-l 参数来禁用
// Goland Run：
//  go test -c gin-sqlx-base/tests/gomonkey
// Goland Debug：
//  go test -c -gcflags "all=-N -l" gin-sqlx-base/tests/gomonkey
func specialHandle() {
	_ = gomonkey.ApplyFunc(Hello, func() string {
		return "haha"
	})
	// defer helloPatch.Reset()
}
