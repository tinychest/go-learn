package util

import (
	"fmt"
	"testing"
)

func TestPrintUtils(t *testing.T) {
	strSlice := []string{"1", "2", "3"}

	fmt.Println(strSlice)
	PrintSliceInfo(strSlice)
}
