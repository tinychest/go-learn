package performance

// go build -gcflags='-m -l' escape_example.go

// 指针逃逸
// 函数返回值如果是地址引用的类型，必然逃逸到堆上（不然，空返回一个已经被回收的地址有什么用...）
// 没有逃逸，不会打印出相关信息

// EscapeMap
// .\escape_example.go:10:26: map[string]string{} escapes to heap:
// .\escape_example.go:10:26:   flow: ~r0 = &{storage for map[string]string{}}:
// .\escape_example.go:10:26:     from map[string]string{} (spill) at .\escape_example.go:10:26
// .\escape_example.go:10:26:     from return map[string]string{} (return) at .\escape_example.go:10:2
// .\escape_example.go:10:26: map[string]string{} escapes to heap
func EscapeMap() map[string]string {
	return map[string]string{}
}

// EscapeSlice
// .\escape_example.go:14:17: []string{} escapes to heap:
// .\escape_example.go:14:17:   flow: ~r0 = &{storage for []string{}}:
// .\escape_example.go:14:17:     from []string{} (spill) at .\escape_example.go:14:17
// .\escape_example.go:14:17:     from return []string{} (return) at .\escape_example.go:14:2
// .\escape_example.go:14:17: []string{} escapes to heap
func EscapeSlice() []string {
	return []string{}
}

func NoEscapeArray() [2]int {
	return [2]int{}
}

// EscapeTooBig
// .\escape_example.go:25:14: make([]int, 10000) escapes to heap:
// .\escape_example.go:25:14:   flow: {heap} = &{storage for make([]int, 10000)}:
// .\escape_example.go:25:14:     from make([]int, 10000) (too large for stack) at .\escape_example.go:25:14
// .\escape_example.go:25:14: make([]int, 10000) escapes to heap
// 那这个限制具体是多大？ （runtime/stack.go）20KB 逃逸 → 8KB 没有逃逸 → 16KB 逃逸 → 10KB 逃逸 → 额 8KB + 1B 逃逸（临界值就是 8KB）
func EscapeTooBig() {
	// 64 位的机器上 int 是 64 位，也就是 8 byte
	// var _ = make([]int, 4*1024*1024)
	var _ = make([]int, 1024*8+1)
}

// EscapeOfUnknownSize
// .\escape_example.go:25:14: make([]int, 10000) escapes to heap:
// .\escape_example.go:25:14:   flow: {heap} = &{storage for make([]int, 10000)}:
// .\escape_example.go:25:14:     from make([]int, 10000) (too large for stack) at .\escape_example.go:25:14
// .\escape_example.go:25:14: make([]int, 10000) escapes to heap
func EscapeOfUnknownSize() {
	size := 1
	var _ = make([]int, size)
}

// EscapeOfUnknownType
// .\escape_example.go:36:9: "" escapes to heap:
// .\escape_example.go:36:9:   flow: ~r0 = &{storage for ""}:
// .\escape_example.go:36:9:     from "" (spill) at .\escape_example.go:36:9
// .\escape_example.go:36:9:     from return "" (return) at .\escape_example.go:36:2
// .\escape_example.go:36:9: "" escapes to heap
func EscapeOfUnknownType() interface{} {
	return ""
}

// EscapeOfSpecialClosure
// .\escape_example.go:66:2: EscapeOfSpecialClosure capturing by ref: n (addr=false assign=true width=8)
// .\escape_example.go:68:9:     from return func literal (return) at .\escape_example.go:68:2
// .\escape_example.go:66:2: n escapes to heap:
// .\escape_example.go:66:2:   flow: {storage for func literal} = &n:
// .\escape_example.go:66:2:     from n (captured by a closure) at .\escape_example.go:69:3
// .\escape_example.go:66:2:     from n (reference) at .\escape_example.go:69:3
// .\escape_example.go:66:2: moved to heap: n
// .\escape_example.go:68:9: func literal escapes to heap
func EscapeOfSpecialClosure() func() {
	n := 0

	return func() {
		n++
	}
}
