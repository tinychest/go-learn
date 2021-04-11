package basic

import (
	"fmt"
	"testing"
)

// 被打上标签的代码是会随着文档代码流线性执行的
func TestGoto(t *testing.T) {
	var flag = false
	if flag {
		goto label
	}

label:
	fmt.Println("走到 Label 了")
}
