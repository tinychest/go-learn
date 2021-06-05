package design_pattern

import (
	"fmt"
	"strings"
	"testing"
)

// 从 go 的学习视频中看到，自己也觉得，并且有良心弹幕说到了，下面的模式在实际开发中非常有用
func TestAModel(t *testing.T) {
	makeFunc := func(suffix string) func(string) string {
		return func(name string) string {
			if strings.HasSuffix(name, suffix) {
				name = name + suffix
			}
			return name
		}
	}

	jpgFunc := makeFunc(".jpg")
	txtFunc := makeFunc(".txt")

	fmt.Println(jpgFunc("abc"))
	fmt.Println(jpgFunc("abc.jpg"))
	fmt.Println(txtFunc("abc"))
	fmt.Println(txtFunc("abc.jpg"))
}
