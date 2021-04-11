package concurrency

import (
	"testing"
	"time"
)

// 现象是 5 秒的任务来不及执行就结束了，因为 main 函数执行的任务结束了
func TestGoRoutineChannel(t *testing.T) {
	// 开启一个 goroutine 执行一个需要5秒的任务
	go WaitFiveSeconds()
	// 主函数中防止一个需要执行3秒的任务
	WaitThreeSeconds()
}

// 成功执行
func MainMethodShowDoRight() {
	// 这样的通道读取都是阻塞的
	ch := make(chan int)

	// 开启一个 goroutine 执行一个需要5秒的任务，并且在执行完毕后，向通道发送消息
	go WaitFiveSecondsSendAndWaitFiveSecondsCloseChannel(ch)

	// 主函数中防止一个需要执行3秒的任务，并且在执行完毕后，等待通道的消息
	// WaitThreeSecondsAndWaitChannelSend(ch)
	WaitThreeSecondsThenWaitChannelClose(ch)
}

func WaitThreeSeconds() {
	time.Sleep(3 * time.Second)
	println("需要 3 秒的任务执行完毕")
}

func WaitFiveSeconds() {
	time.Sleep(5 * time.Second)
	println("需要 5 秒的任务执行完毕")
}

func WaitThreeSecondsAndWaitChannelSend(ch chan int) {
	time.Sleep(3 * time.Second)
	println("需要 3 秒的任务执行完毕，开始等待通道消息")
	// 从通道取到值就结束
	<-ch
}

func WaitThreeSecondsThenWaitChannelClose(ch chan int) {
	time.Sleep(3 * time.Second)
	println("需要 3 秒的任务执行完毕，开始等待通道消息")
	// 改成从信道取值的 for-Each 的话，则信道关闭才认为结束
	for value := range ch {
		println("从通道收到的值为: ", value)
	}
}

func WaitFiveSecondsSendAndWaitFiveSecondsCloseChannel(ch chan int) {
	time.Sleep(5 * time.Second)
	println("需要 5 秒的任务执行完毕，向通道发送一条数据")
	ch <- 1

	time.Sleep(5 * time.Second)
	println("又 5 秒的任务执行完毕，关闭 Channel - 告诉 main 方法任务完成，可以结束了")
	close(ch)
}
