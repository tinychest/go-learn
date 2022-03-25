package regex

import (
	"regexp"
	"testing"
)

// ^，中文名脱字符，在正则表达式中，除了表达开头的概念，还可以表示 非 的概念，详见下例

func TestTuo(t *testing.T) {
	// 在 [] 中，表示目标串中必须含有不存在于下面集合中的字符
	r := `[^abcd]`
	s := `f`

	t.Log(regexp.MustCompile(r).MatchString(s))
}
