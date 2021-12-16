package regex

import (
	"go-learn/util"
	"regexp"
	"testing"
)

func TestSplit(t *testing.T) {
	data := regexp.MustCompile(`\s+`).Split("a > 1", -1)
	util.PrintSlice(data)

	data = regexp.MustCompile(`\s+`).Split("a > 1", 1)
	util.PrintSlice(data)
}
