package slice

import "testing"

func TestGeneric(t *testing.T) {
	stringSlice := New[string](1)
	stringSlice.Append("1")
	stringSlice.Append("2")
	stringSlice.Append("3")
	stringSlice.Append("3")
	t.Log(stringSlice)

	stringSlice.DeleteByIndex(3)
	t.Log(stringSlice)
}
