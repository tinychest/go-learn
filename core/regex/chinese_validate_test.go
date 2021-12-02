package regex

import (
	"fmt"
	"regexp"
	"testing"
)

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

	fmt.Println(reg1.MatchString("一")) // true
	fmt.Println(reg3.MatchString("一")) // false
}

// 测试什么样的字符串才能匹配 demoRegex3
// 结论：只要是中括号中的任意字符都行
func TestDemoRegex3Match(t *testing.T) {
	var reg3 = regexp.MustCompile(demoRegex3)

	fmt.Println(reg3.MatchString("一"))
	fmt.Println(reg3.MatchString(``))
}
