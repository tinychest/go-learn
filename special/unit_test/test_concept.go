package unit

import (
	"testing"
)

// 测试方法要求：func TestXxx(xxx *testing.T) {}（名以 Test 开头，参数类型为 *testing.T）
// 测试文件名要求：xxx_test.go（以 _test.go 结尾）

// 测试文件：go test xxx_text.go xxx.go
// 测试方法：go test -test.run TestXxx | go test -test.run=TestXxx（忘记了怎么办，直接在 Golang UI 中运行一个测试方法，然后查看控制台最上面的详细命令）

// -v 打印详细信息（比如，日志）

func TestXxx(t *testing.T) {
	println("I can not run because file name is not xxx_test style...")
}
