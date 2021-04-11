package basic

import (
	"strconv"
	"testing"
)

func TestCast(t *testing.T) {
	// string → int
	if atoi, err := strconv.Atoi("123"); err != nil {
		println("string → int 出错: ", err)
	} else {
		println(atoi)
	}

	// string → int64
	if intValue, err := strconv.ParseInt("123", 10, 64); err != nil {
		println("string → int64 出错", intValue)
	} else {
		println(intValue)
	}

	// int → string
	stringValue := strconv.Itoa(123)
	println(stringValue)

	// int64 → string
	stringValue = strconv.FormatInt(123, 10)
	println(stringValue)
}
