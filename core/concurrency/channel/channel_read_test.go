package channel

import (
	"fmt"
	"testing"
)

// 不带缓存的通道
func TestChannelRead(t *testing.T) {
	c := make(chan int)

	// select（非阻塞，注意，加了 default 才是非阻塞）
	// select {
	// case v, ok := <-c:
	// 	fmt.Printf("value：%d，ok：%t\n", v, ok)
	// default:
	//
	// }

	// 方式一（阻塞）
	<-c

	// 方式二（阻塞），ok 代表通道是否关闭，因为关闭的通道也可以读取出值
	v, ok := <-c
	fmt.Printf("value：%d，ok：%t\n", v, ok)
}

// 带有缓存的通道：make(chan int, n)
// 通道缓冲区大小默认是 0，即没有缓冲区：make(chan int)、make(chan int, 0)
// 读带有缓存通道的结束时机：chan close and chan is empty
func TestBufferedChannelRead(t *testing.T) {
	bc := make(chan int, 4)
	bc <- 1
	bc <- 2
	bc <- 3
	close(bc)

	// 方式一
	// break 需要指定标签退出
	// select：选择 case 中任意一个达成条件的执行，没有达成条件的就一直等待
forLoop:
	for {
		v, ok := <-bc
		if !ok && len(bc) == 0 {
			break forLoop
		}
		t.Logf("从通道读取：%d\n", v)
	}

	// 方式二（推荐）
	for i := range bc {
		t.Logf("从通道读取：%d\n", i)
	}
}
