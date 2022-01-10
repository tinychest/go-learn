package todo

import (
	"fmt"
	"testing"
	"unsafe"
)

// https://philpearl.github.io/post/anathema/
func TestRun(t *testing.T) {
	var limited int
	var out int
	doThing(out, func(out interface{}) {
		if limited == 0 {
			fmt.Printf("limited is zero. %d\n", limited) // Prints 42
		}
		limited++
	})
}

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}

//go:noinline
func doThing(out interface{}, f func(out interface{})) {
	p := (*eface)(unsafe.Pointer(&out)).data
	*(*int)(p) = 42
	f(out)
}
