package basic

import (
	"strings"
	"testing"
)

func TestStringReplace(t *testing.T) {
	strReplacer := strings.NewReplacer("1", "2", "2", "3", "3", "4")

	repStr := strReplacer.Replace("123")

	println(repStr)
}
