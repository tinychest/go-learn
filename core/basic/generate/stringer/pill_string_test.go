package stringer

import (
	"fmt"
	"testing"
)

// 测试通过 stringer 命令为 Pill 常量生成的实现的 string 方法
func TestPillString(t *testing.T) {
	// 没有生成时：1 运行异常 1
	fmt.Println(Aspirin)
	fmt.Printf("%s\n", Aspirin)
	fmt.Printf("%v\n", Aspirin)

	// 生成后：Aspirin Aspirin Aspirin
}
