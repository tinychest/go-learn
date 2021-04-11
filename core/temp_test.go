package core

import (
	"fmt"
	"strings"
	"testing"
)

/*
<后置取>
个数理解：取几个字符
下标理解：截至（不包括） → 3 = 0 1 2
<前置取>
个数理解：从第几个字符开始（不包括）
下标理解：从这里开始（包括）
*/
func TestLimit(t *testing.T) {
	phone := "18942350318"
	phone = phone[:3] + "****" + phone[7:]
	fmt.Println(phone)
}

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
