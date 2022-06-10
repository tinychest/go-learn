package gomonkey

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"testing"
)

// 递归了

func p() {
	fmt.Println("2")
}

func TestAOP(t *testing.T) {
	patch := gomonkey.ApplyFunc(p, func() {
		fmt.Println("1")
		p()
		fmt.Println("3")
	})
	defer patch.Reset()

	p()
}
