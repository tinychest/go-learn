package string

import (
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"
)

// - string 不能为 nil，string 是不可变的
// - string 是可以比较的（按照字典顺序）
// - 字符串编码为 UTF-8，使用 1-4 字节表示一个字符（汉字 3 个字节，emoji 表情 4 个字节），一个 rune 代表一个 unicode 字符

// 说到 string 就离不开字符，string 可以理解为字符切片（因不可变性，理解为字符数组可能更恰当一些）
func TestString(t *testing.T) {
	// var str = "123"
	// str = str[:len(str) - 1]
}

// 获取 string 的字符数
// string 直接通过下标取值得到的类型是 uint8 byte 字节，是对应的字节数组的长度，而不是对应的字符数组的长度
func TestLen(t *testing.T) {
	printLen := func(str string) {
		fmt.Println(len(str))                    // 字节数
		fmt.Println(len([]byte(str)))            // 字节数
		fmt.Println(len([]rune(str)))            // 字符数
		fmt.Println(utf8.RuneCountInString(str)) // 字符数
	}

	printLen("我") // 3 3 1 1
	printLen("😃") // 4 4 1 1
}

// 获取 string 指定下标的字符
func TestCharAt(*testing.T) {
	s := "我是"
	fmt.Println(strconv.QuoteRune([]rune(s)[0]))
	fmt.Printf("%c\n", []rune(s)[0])
}

// 遍历 string 的每一个字符
// strconv.QuoteRune 根据二进制数值，查询 UTF-8 码表获取对应的字符
func TestTraversingChar(*testing.T) {
	str := "Golang够浪😊"

	// 方式 1（value 是 int32 rune 类型）
	for _, value := range str {
		fmt.Print(strconv.QuoteRune(value) + " ")
	}
	fmt.Println()

	// 方式 2
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)
		str = str[size:]
	}
}
