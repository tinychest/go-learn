package _defer

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	var n = 0

	defer func(a int) {
		fmt.Println(a)
	}(n) // 此时就已经固定了 n 的参数值为 0

	n = 1
}
