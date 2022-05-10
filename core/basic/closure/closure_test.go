package closure

import "testing"

func TestClosure(t *testing.T) {
	outModifyTest(t)
	inModifyTest(t)
	inPassModifyTest(t)
}

func outModifyTest(t *testing.T) {
	n := 0
	f := func() { t.Log(n) }
	n++

	f()
}

func inModifyTest(t *testing.T) {
	n := 0
	f := func() { n++ }

	f()
	t.Log(n)
}

func inPassModifyTest(t *testing.T) {
	n := 0
	f := func() { n++ }

	whatever(f)
	t.Log(n)
}

func whatever(f func()) {
	f()
}
