package convey_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 简单介绍一下，goconvey 是 Go 的第三方测试框架（https://github.com/smartystreets/goconvey）
// 1、比 Go 原生自带的 testing 包具有更简单易用的语法（断言：So、testify 的 assert(返回 bool 值) 和 require(会中断)）
// 2、在项目目录下，启动运行《Go convey 编译生成的二进制文件》，访问 8080 端口能够在浏览器中查看项目中所有单元测试的结果（只要当前程序运行，浏览器的结果就会实时更新）
// 3、还支持 Python 脚本，来实现自动化测试
func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				// 通过（看不太懂）
				// .
				// 1 total assertion

				// 不通过：详细原因
				So(x, ShouldEqual, 2)
			})
		})
	})
}
