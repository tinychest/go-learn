package regex

import (
	"fmt"
	"regexp"
	"testing"
)

// 从目标串中找出所有，不以 a 打头，WWW 结尾 的单词
func TestCase01(t *testing.T) {
	r := `\b[^a ]*WWW\b`
	reg := regexp.MustCompile(r)

	// Go 中不支持 ?<
	// r := `\S*(?<!a)WWW`
	// reg := regexp.MustCompile(r2)

	s := " WWW aWWW bWWW cWWW 123WWW abcWWW"
	res := reg.FindAllString(s, -1)
	t.Log(res)
}

// 目标串是否是 以多个端口号通过 || 隔开的格式
func TestRegPortSplit(t *testing.T) {
	// 1-65535
	// portRange := `^([0-9]|[1-9]\d|[1-9]\d{2}|[1-9]\d{3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`
	// reg := regexp.MustCompile(portRange)
	// for i := 1; i <= 65536; i++ {
	// 	ok := reg.MatchString(strconv.Itoa(i))
	// 	if !ok {
	// 		t.Log(i)
	// 		break
	// 	}
	// }

	portRange := `([0-9]|[1-9]\d|[1-9]\d{2}|[1-9]\d{3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])`
	regStr := fmt.Sprintf(`^%s(\|\|%s)?$`, portRange, portRange)
	reg := regexp.MustCompile(regStr)
	t.Log(reg.MatchString("8080||5566||00"))
}