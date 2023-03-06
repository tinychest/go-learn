package performance

import "unsafe"

// 来自：https://mp.weixin.qq.com/s/RNHpyiiOP5YLIb3UHqtcrQ

// 查看逃逸分析结果：go build -gcflags=-m escape_analyse_example.go

// 经典的逃逸
func escapeFunc() (*string, *string) {
	s1 := new(string)
	s2 := new(string)
	return s1, s2
}

// 皮一下，通过 safe 整一个 noescape 的方法
func noescapeFunc() *string {
	s := new(string)
	return (*string)(noescape(unsafe.Pointer(s)))
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}
