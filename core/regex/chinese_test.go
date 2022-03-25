package regex

import (
	"regexp"
	"testing"
)

// 正则表达式中，表达字符串的内容是一个中文汉字：^[\u4E00-\u9FA5]$
// 约定 \u 开头的是一个 Unicode 码的字符。范围在 '\u0000' 到 '\uFFFF' 之间
//
// - u 在正则表达式中不是特殊字符，所以使用转义符去使其表达出特殊含义是非法的
// - 正则表达式中，使用 \ 去转译 u ，只能通过 "" 去修饰，让 \ 在 Go 源码中表达出含义

// 判断一个字符是否是中文汉字
var demoRegex1 = "^[\u4E00-\u9FA5]$"

// 正则表达式非法，通过 regexp.MustCompile 会得到 panic
var demoRegex2 = `^[\u4E00-\u9FA5]$`

// 正则表达式合法，但是已经失去了表达中文汉字的含义
var demoRegex3 = `^[\\u4E00-\\u9FA5]$`

// 测试上面三个正则表达式表达的含义
// 结论：已经在表达式的头上了
func TestChineseRegexMatch(t *testing.T) {
	var reg1 = regexp.MustCompile(demoRegex1)
	// _ = regexp.MustCompile(demoRegex2)
	var reg3 = regexp.MustCompile(demoRegex3)

	t.Log(reg1.MatchString("一")) // true
	t.Log(reg3.MatchString("一")) // false
}

// 测试什么样的字符串才能匹配 demoRegex3
// 结论：只要是中括号中的任意字符都行
func TestDemoRegex3Match(t *testing.T) {
	var reg3 = regexp.MustCompile(demoRegex3)

	t.Log(reg3.MatchString("一"))
	t.Log(reg3.MatchString(``))
}
