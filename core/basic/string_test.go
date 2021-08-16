package basic

import (
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"
)

// Go 中，string 不能为 nil，string 是不可变的

// 说到 string 就离不开字符，string 可以理解为字符切片（因不可变性，理解为字符数组可能更恰当一些）
func TestString(t *testing.T) {
	// 1、支持同 切片 数组的范围取值转化
	// var str = "123"
	// str = str[:len(str) - 1]

	// 2、string 是可以比较的（按照字典顺序）
}

// 因为 Go 语言中的字符串编码为 UTF-8，使用 1-4 字节就可以表示一个字符，一个 rune 代表一个 unicode 字符

// unicode/utf8.RuneCountInString
func TestLen(t *testing.T) {
	str := "Golang够浪"
	fmt.Println(len(str))                    // 12
	fmt.Println(len([]rune(str)))            // 8
	fmt.Println(utf8.RuneCountInString(str)) // 8

	// string 通过下标取值得到的类型是 uint8(byte)，所以 string 可以转化成：[]byte

	// 类型是：int32(rune)
	for _, value := range str {
		// 根据 ASCII 码获取对应的字符
		fmt.Print(strconv.QuoteRune(value) + " ")
	}
	fmt.Println()
}
