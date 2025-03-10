package main

import (
	"runtime"
)

// 这里得到的结果是 1，在 Test 中得到的结果是 2
func main() {
	println(runtime.NumGoroutine())
}
