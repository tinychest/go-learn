package regex

import (
	"go-learn/tool"
	"regexp"
	"testing"
)

func TestSplit(t *testing.T) {
	data := regexp.MustCompile(`\s+`).Split("a > 1", -1)
	tool.PrintSlice(data)

	data = regexp.MustCompile(`\s+`).Split("a > 1", 1)
	tool.PrintSlice(data)
}
