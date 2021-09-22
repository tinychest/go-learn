package case_delay

import (
	"fmt"
	"testing"
)

func TestDelay(*testing.T) {
	case1()
}

func case1() {
	n := 0
	f := func() {fmt.Println(n)}
	n++

	f()
}
