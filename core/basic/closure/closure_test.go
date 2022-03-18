package closure

import "testing"

func TestClosure(t *testing.T) {
	n := 0
	f := func() { t.Log(n) }
	n++

	f()
}
