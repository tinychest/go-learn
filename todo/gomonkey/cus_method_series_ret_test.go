package gomonkey

import (
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 定义方法 n 次调用的返回值

func Str() string {
	return "a"
}

func TestDefineRet(t *testing.T) {
	Convey("", t, func() {
		want := []gomonkey.OutputCell{
			{Values: gomonkey.Params{"1"}},
			{Values: gomonkey.Params{"2"}},
			{Values: gomonkey.Params{"3"}},
		}
		str := gomonkey.ApplyFuncSeq(Str, want)
		defer str.Reset()

		Convey("Then Test ApplyFuncSeq Patch", func() {
			So(Str(), ShouldEqual, "1")
			So(Str(), ShouldEqual, "2")
			So(Str(), ShouldEqual, "3")
		})
	})
}
