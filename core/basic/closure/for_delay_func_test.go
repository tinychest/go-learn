package closure

import (
	"testing"
)

func TestDelayFunc(t *testing.T) {
	var fs []func()
	for i := 0; i < 3; i++ {
		fs = append(fs, func() {
			t.Log(i)
		})
	}
	for _, f := range fs {
		f()
	}
}

func TestDelayFuncFix(t *testing.T) {
	var fs []func()
	for i := 0; i < 3; i++ {
		i := i
		fs = append(fs, func() {
			t.Log(i)
		})
	}
	for _, f := range fs {
		f()
	}
}
