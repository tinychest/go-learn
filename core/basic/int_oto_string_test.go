package basic

import (
	"fmt"
	"go-learn/util"
	"reflect"
	"strconv"
	"testing"
)

// atoi：a to i（integer 整型）
// itoa：i to a（array 字符串）
// 核心包 strconv
func TestStringOtoInt(t *testing.T) {
	StringToInt("1")
	sti := StringToInt64("1")
	IntToString(1)
	Int64ToString(1)

	println(sti)

	// numberStr := SumToString(2)
	// println(numberStr)
}

func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func StringToInt64(str string) int64 {
	// base：str 表示的数 的进制，一般都是 10
	//     取值范围：0（当为0时，通过 str 的数字前缀格式来区分，如 0b=二进制、0|0o=八进制、0x=十六进制、无=十进制）
	//     取值范围：[2, 36]（最大值为 36，是因为 10 + 26 个英文字母）
	//     超过范围：err invalid base
	// bitSize：表示需要多少位来存储该数，一般都是 64
	//     取值范围：[0, 64]
	//     超过范围：err invalid bit size
	i, err := strconv.ParseInt(str, 10, -1)
	if err != nil {
		panic(err)
	}
	return i

	// 1、bitSize 例1
	// 将 16 进制的 -354634382 转化位对应的 10 进制数为 -4294967295
	// 其中 4294967295 超过了 32-1=31 位能表示的最大数 2147483648
	// v32 := -354634382
	// if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
	//     // 不会执行
	//     fmt.Printf("%T, %v\n", s, s)
	// }
	// if s, err := strconv.ParseInt(v32, 16, 64); err == nil {
	//     // 会执行
	//     fmt.Printf("%T, %v\n", s, s)
	// }

	// 2、bitSize 例2
	// 假如 bitSize 为1，任何负数 → -1，可以接受 0，存储不了任何正数
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func SumToString(number interface{}) string {
	paramValue := reflect.ValueOf(number)
	util.PtrUnReference(&paramValue)

	switch paramValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	default:
		panic("注意方法名！你传个非整型干嘛，瞧不起我 Go 没有泛型么？🐕")
	}

	return fmt.Sprint(number)
}
