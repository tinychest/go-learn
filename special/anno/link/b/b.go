package b

import (
	_ "unsafe"
)

//go:linkname b go-learn/special/link/a.A
func b() {
	println("啊哈，我是 b")
}
