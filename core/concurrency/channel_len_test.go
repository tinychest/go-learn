package concurrency

import (
	"testing"
)

// 通道在作为参数传给 len 和 cap 得到的结果和 数组/切片 作为参数的效果是一样的
// len：通道中的元素个数
// cap：通道的缓存大小
func TestChannelLen(t *testing.T) {
	intChannel := make(chan int, 2)
	println(cap(intChannel))
	println(len(intChannel))

	intChannel <- 1
	println(cap(intChannel))
	println(len(intChannel))
}
