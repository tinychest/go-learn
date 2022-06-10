package tool

import (
	"testing"
)

func TestPrintUtils(t *testing.T) {
	strSlice := []string{"1", "2", "3"}

	t.Log(strSlice)
	PrintSlice(strSlice)
}
