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
	fmt.Println(res)
}
