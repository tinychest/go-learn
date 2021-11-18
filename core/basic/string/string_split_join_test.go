package string

import (
	"fmt"
	"go-learn/util"
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	fmt.Println(strings.Join(nil, ",") == "")
}

// 注意：返回的并不是 null，而是 len:1 cap:1 仅含一个空串元素的切片
func TestSplit(t *testing.T) {
	splits := strings.Split("", ",")

	util.PrintSlice(splits)
	fmt.Println(splits[0] == "")
}
