package regex

import (
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	case1(t)
	case2(t)
}

// 最简单的样例
func case1(t *testing.T) {
	r := `^aaa.*bbb$`
	ys := `aaa___bbb`
	ns := `aac___bbb`

	t.Log(regexp.MatchString(r, ys))
	t.Log(regexp.MatchString(r, ns))
}

// 测试 |
func case2(t *testing.T) {
	t.Log(regexp.MatchString(`^|$`, "a"))
	t.Log(regexp.MatchString(`^ab|cd$`, ""))     // 等同于 (^ab)|(cd$)
	t.Log(regexp.MatchString(`^(ab)|(cd)$`, "")) // 表示 ab | cd
	t.Log(regexp.MatchString(`^[ab|cd]$`, "a"))  // 表示 a | b | c | | | d
}
