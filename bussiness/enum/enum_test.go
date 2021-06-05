package enum

import "testing"

func TestEnumCast(t *testing.T) {
	println(Parse(""))
	println(Parse("123"))
	println(Parse("MALE"))
}
