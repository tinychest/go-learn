package regex

import (
	"regexp"
	"testing"
)

// 替换所有独立的单词 hello

func TestReplace(t *testing.T) {
	s := `hello helloworld hello world hello  world`
	res := regexp.MustCompile(`\bhello\b`).ReplaceAllString(s, "fuck")
	t.Log(res)
}
