package set

import "testing"

func TestGeneric(t *testing.T) {
	// 方式一
	stringSet := New[string](1)
	stringSet.Add("1")
	stringSet.Add("2")
	stringSet.Add("3")
	stringSet.Add("3")
	t.Log(stringSet.ToSlice())

	// 方式二
	s := NewBy[string]([]string{"1", "2", "3"}).ToSlice()
	t.Log(s)
}
