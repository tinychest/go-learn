package case_delay

import (
	"testing"
)

func TestDelay(t *testing.T) {
	case1(t)
}

func case1(t *testing.T) {
	n := 0
	f := func() { t.Log(n) }
	n++

	f()
}
