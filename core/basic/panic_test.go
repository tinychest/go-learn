package basic

import (
	"fmt"
	"testing"
)

// func panic(v interface{})
// panic 的参数类型并不是 error
// func recover() interface{}
// recover 的返回值类型也不是 error

func TestPanic(t *testing.T) {
	panicTwiceTest(t)
}

func panicTwiceTest(t *testing.T) {
	defer func() {
		// 只能捕获到第二次 panic 的信息
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 第二次 panic
	defer panic(456)
	// 第一次 panic
	panic(123)
}
