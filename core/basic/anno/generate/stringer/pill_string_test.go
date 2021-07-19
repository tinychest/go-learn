package stringer

import (
	"fmt"
	"testing"
)

// 测试通过 stringer 命令为 Pill 常量生成的 string 方法实现
func TestPillString(t *testing.T) {
	fmt.Println(Aspirin)
	fmt.Printf("%s\n", Aspirin)
	fmt.Printf("%v\n", Aspirin)

	// 未生成 pill_string.go：1 运行异常 1
	// 生成后：Aspirin Aspirin Aspirin
}
