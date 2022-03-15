package string

import (
	"strings"
	"testing"
)

func TestStringReplace(t *testing.T) {
	res := strings.NewReplacer("a", "A", "b", "B", "c", "C").Replace("abc")

	t.Log(res)
}
