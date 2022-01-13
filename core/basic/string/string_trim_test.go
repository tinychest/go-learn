package string

import (
	"strings"
	"testing"
)

// HasPrefix 和 TrimPrefix
// TrimLeft：主串中逐字符在子串匹配，然后删除，直至没有在子串中匹配的字符
func TestStringTrim(t *testing.T) {
	s := "User.Univ"
	cutset := "User."

	t.Log(strings.TrimPrefix(s, cutset))
	t.Log(strings.TrimLeft(s, cutset))
}
