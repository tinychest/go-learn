package regex

import (
	"regexp"
	"testing"
)

// 我想描述这样这样一个字符串，3 位数字重复 2 次以上的字符串
// 如：123123123、101101...
// 该如何书写正则表达式，这时就要借助捕获组引用了（BackReferences），但是 Go 标准库好像不支持
// 详见 StackOverflow：https://stackoverflow.com/questions/23968992/how-to-match-a-regex-with-backreference-in-go
// 一些大佬开发的三方库支持：https://golangbyexample.com/golang-regex-backreferences/
func TestBackReferences(t *testing.T) {
	var r string

	s := `123123123`
	// 组名引用（好像不支持） → GG
	r = `(?P<dup>\d{3})\P{dup}+`

	// 下标引用1（编译器已经提示：This named group reference syntax is not supported in this regex dialect） → GG
	r = `(\d{2})(?P=1)`

	// 下标引用2（GoRegExp 语法检查通过，但是运行时 panic） → GG
	r = `^(\d{2})\1+$`
	reg := regexp.MustCompile(r)
	t.Log(reg.MatchString(s))
}
