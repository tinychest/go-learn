package string

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"
)

// - string 不能为 nil，string 是不可变的（通过下标访问字符串指定位的字节，是不可寻址的，即字节是不可变的）
// - string 是可以比较的（按照字典顺序）
// - 字符串编码为 UTF-8，使用 1-4 字节表示一个字符（汉字 3 个字节，emoji 表情 4 个字节），一个 rune 代表一个 unicode 字符
// - for range []byte(string) 不会做深拷贝，同时这也是高效遍历字节的方式
// - 字符串当作切片使，是一种语法糖（本质是先转成字节切片）

// 一些常见方法的说明：
// strconv.QuoteRune 返回将字符使用单引号引起来的字符串结果
// bytes.Runes []byte → []rune
// utf8.RuneCount 字符串（字节数组）中的字符数
// utf8.RuneCountInString 字符串中的字符数
// utf8.DecodeRune 辅助遍历字符串（字节数组）中的字符
// utf8.DecodeRuneInString 辅助遍历字符串中的字符

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

// 演示 strconv.QuoteRune 方法
func TestCharAt(*testing.T) {
	s := "我是"
	fmt.Printf("%c\n", []rune(s)[0])
	fmt.Printf("%s\n", strconv.QuoteRune([]rune(s)[0]))
}

// 遍历 string 的每一个字符
// - for range 直接遍历
// - string → []rune → 下标遍历
// - 通过 bytes 包的相关方法（[]byte → []rune，详见源码） → 下标遍历
// - 通过 utf8 包的相关方法，遍历到每一个字符
func TestTraversingChar(t *testing.T) {
	str := "Golang够浪😊"

	// 方式 1（value 是 int32 rune 类型）
	t.Log("--- 遍历方式一 ---")
	for _, value := range str {
		fmt.Print(strconv.QuoteRune(value) + " ")
	}
	fmt.Println()

	// 方式 2
	t.Log("--- 遍历方式二 ---")
	rs := []rune(str)
	for i := 0; i < len(rs); i++ {
		fmt.Print(strconv.QuoteRune(rs[i]) + " ")
	}
	fmt.Println()

	// 方式 3（多此一举的感觉）
	t.Log("--- 遍历方式三 ---")
	rs = bytes.Runes([]byte(str))

	for i := 0; i < len(rs); i++ {
		fmt.Print(strconv.QuoteRune(rs[i]) + " ")
	}
	fmt.Println()

	// 方式 4
	t.Log("--- 遍历方式四 ---")
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)
		str = str[size:]
	}
}
