package convey_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
【简介】
goconvey 是 Go 的第三方测试框架（https://github.com/smartystreets/goconvey）
- 比 Go 原生自带的 testing 包具有更简单易用的语法（断言：So、testify 的 assert(返回 bool 值) 和 require(会中断)）
- 在项目目录下，启动运行 "go convey 编译生成的二进制文件"，访问 8080 端口能够在浏览器中查看项目中所有单元测试的结果（只要当前程序运行，浏览器的结果就会实时更新）

【安装】
go get github.com/smartystreets/goconvey

【简单样例】
以下面的三方简单方法作为测试对象，编写的 convey 用例如下
只要在当前目录下执行 "goconvey" 就能通过 "http://localhost:8080" 看到实时的结果

【其他】
在 convey 中，如果想表达流程（第一步走通了，然后，走第二步，这样），convey 应该是嵌套形式，而不是并列形式
*/

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func TestAdd(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("test add", t, func() {
		So(add(1, 1), ShouldEqual, 2)
	})
}

func TestSubtract(t *testing.T) {
	Convey("test subtract", t, func() {
		So(subtract(1, 1), ShouldEqual, 0)
	})
}

func TestMultiply(t *testing.T) {
	Convey("test multiply", t, func() {
		So(multiply(1, 1), ShouldEqual, 1)
	})
}
