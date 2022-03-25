package regex

import (
	"regexp"
	"testing"
)

// 演示什么是回溯
func TestLookup(t *testing.T) {
	s := `aaab`

	// a{1,3} 一来直接匹配 3 个 a，然后发现后面一个 a 无法匹配上目标串的 b，只能回溯，匹配 2 个 a
	// 这就是回溯，所以假如 Go 支持独占模式，r3 匹配结果将是 false
	r1 := `^a{1,3}ab$`
	r2 := `^a{1,3}?ab$`
	r3 := `^a{1,3}+ab$`

	t.Log(regexp.MatchString(r1, s)) // true
	t.Log(regexp.MatchString(r2, s)) // true
	t.Log(regexp.MatchString(r3, s)) // error parsing regexp: invalid nested repetition operator: `{1,3}+`
}
