package channel

import (
	"testing"
)

func TestChannel(t *testing.T) {
	// 通过 new 关键字创建的 channel 起始就是 nil
	// c := new(chan int)
	//	if *c == nil {
	//		t.Log("c is nil")
	//	}

	// nil channel
	var nc chan int
	close(nc) // 关：panic: close of nil channel
	nc <- 1   // 写：阻塞
	<-nc      // 读：阻塞

	// closed channel
	cc := make(chan int)
	close(cc)
	close(cc) // 关：panic: close of nil channel
	cc <- 1   // 写：panic: send on closed channel
	<-cc      // 读：OK

	// 结论，可以从已经关闭的通道读取出类型零值
}
