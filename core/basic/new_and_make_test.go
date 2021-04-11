package basic

import (
	"testing"
)

// 切片、映射、通道（经测试，不用 make 构建的通道无法使用），使用make
// 数组、结构体、其他的所有值类型，使用new
func TestNewAndMake(t *testing.T) {
	type example struct{}

	_ = make([]int, 0)
	_ = make(map[string]string, 0)
	_ = make(chan int, 0)

	_ = new([]int)
	_ = new(map[string]string)
	_ = new(chan int)

	_ = new([0]int)
	_ = new(example)
	_ = new(int)

	// abnormal
	// _ = make([0]int, 0)
	// _ = make(example, 0)
	// _ = make(int, 0)
}

// new 和 make 在通道类型上的表现来说
func forChannelTest(t *testing.T) {
	// 官网：The value of an uninitialized channel is nil
	// 只有通过 make 得到的通道才能算是初始化了的“initialized”，另外两个的表现就像是向通道发送值发送失败
	// var intChannel chan int
	// var intChannel chan int = nil
	// var intChannelPtr = new(chan int) 和 var intChannel = *intChannelPtr
	intChannel := make(chan int)

	go func() {
		intChannel <- 1
	}()

	println("值为：", <-intChannel)
}
