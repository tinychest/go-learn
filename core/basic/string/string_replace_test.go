package string

import (
	"strings"
	"testing"
)

func TestStringReplace(t *testing.T) {
	strReplacer := strings.NewReplacer("a", "A", "b", "B", "c", "C")

	repStr := strReplacer.Replace("abc")

	println(repStr)
}
