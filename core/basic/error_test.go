package basic

import (
	"fmt"
	"testing"
)

// GO 语言目前无法避免每个方法调用后的异常处理，所有需要定好规则，发生错误时，需要对如下几个方面做好约定
// 1、打印程序错误日志的地方：哪里出错哪里打印，因为哪里出错，哪里对出错的场景最清楚，且目前 GO 无法做到优雅的判断异常
// 2、区分前端展现的错误信息和程序日志打印的错误信息
func TestError(t *testing.T) {
	println("外层开始 === {")

	defer func() {
		println("外层 defer 1 - recover之前")
		if err := recover(); err != nil {
			fmt.Printf("外层 defer 1 - %#v\n", err)
		} else {
			println("外层 defer 1 - 不需要 recover")
		}
		println("外层 defer 1 - recover之后")
	}()

	println("调用内层方法之前")
	f()
	println("调用内层方法之后")

	defer func() {
		println("外层 defer 2")
	}()

	println("外层结束 } ===")
}

func f() {
	println("内层开始 === {")

	defer func() {
		println("内层 defer 1")
	}()

	defer func() {
		println("内层 defer 2 - recover之前")
		if err := recover(); err != nil {
			fmt.Printf("%v - %#v\n", "内层 defer 2", err)
		}
		println("内层 defer 2 - recover之后")
	}()

	defer func() {
		println("内层 defer 3")
	}()

	// 抛出异常
	println("异常信息")

	//
	defer func() {
		println("内层 defer 4")
	}()

	// 之后的语句不会再执行了
	println("内层结束 } ===")
}
