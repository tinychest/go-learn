package stringer

import (
	"testing"
)

// 测试通过 stringer 命令为 Pill 常量生成的 string 方法实现
func TestPillString(t *testing.T) {
	t.Log(Aspirin)
	t.Logf("%s\n", Aspirin)
	t.Logf("%v\n", Aspirin)

	// 未生成 pill_string.go：1 运行异常 1
	// 生成后：Aspirin Aspirin Aspirin
}
