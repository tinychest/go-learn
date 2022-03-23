package enum

import "testing"

func TestEnumCast(t *testing.T) {
	t.Log(Parse(""))
	t.Log(Parse("123"))
	t.Log(Parse("MALE"))
}
