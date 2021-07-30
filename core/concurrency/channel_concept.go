package concurrency

import (
	"fmt"
	"testing"
)

// chan xxx：  定义既能接收也能发送的通道类型
// <-chan xxx：定义只能接收的通道类型
// chan<- xxx：定义只能发送的通道类型

// len(chan)：通道中的元素个数
// cap(chan)：通道的缓冲区大小
// close(chan<-)：close 方法不能关闭一个只读通道，这里边隐含着两个意思
//   一个，不能关闭，即不应该有，即创建一个只读通道是不合理的，无法写，也读不出东西
//   二个，因为 chan 类型 可以被 <-chan 或 chan<- 类型接收，关闭通道逻辑应该放在面向对象模型定义中的发送者中

/*
《通道关闭原则》
一个常用的使用Go通道的原则是不要在数据接收方或者在有多个发送者的情况下关闭通道。换句话说，我们只应该让一个通道唯一的发送者关闭此通道
遵从上面的原则，在多个发送者定义的通用逻辑中：

实战：https://gfw.go101.org/article/channel-closing.html
一、1 sender - n receiver：close(chan) - for range chan
二、n sender - 1 receiver：所有引用通道的 Goroutine 都结束了，这个通道会被作为垃圾回收掉，需要额外创建一个通道，这样才不会违反原则
  sender   当符合条件时，直接关闭额外通道
  receiver 从额外通道读，读到就退出
三、m sender - n receiver：思想同上，但是说将关闭额外通道的逻辑无论是放到 sender 还是 receiver 都会违反原则，所以就再起一个 Goroutine 作为中间者，负责关闭额外通道
  sender   当符合条件时，向中间者发出停止信号；从额外通道读，读到就退出
  receiver 当符合条件时，向中间者发出停止信号；从额外通道读，读到就退出
四、1 sender - n receiver - 第三方发起停止：closing 发起关闭信号的通道、closed 辅助关闭的通道、data 数据通道
  第三方从发起关闭信号，如果阻塞从辅助通道读一个值
  sender   读到关闭信号，退出，关闭 辅助通道 和 数据通道
  receiver for range 数据通道
*/

func TestChannel(t *testing.T) {
	var intChannel = make(chan int)
	close(intChannel)

	// 未关闭 - 阻塞住，拿到值才会返回
	// 已关闭 - 非阻塞，得到通道存值类型对应的零值 和 false（从一个已经关闭的通道接收数据，将直接得到通道类型对应的 零值 - 这一条相当重要，且实用，需要记住！！！）
	// （还可以通过 for range chan 的语法判断通道有没有关闭）
	value, ok := <-intChannel
	fmt.Printf("value：%d，ok：%t\n", value, ok)

	// 不能关闭 nil 通道：var nilChannel chan int; close(nilChannel) // panic: close of nil channel
	// 向 nil 通道发送或者接收数据，都会被永久阻塞

	// 不能关闭已经关闭的通道：close(intChannel)                       // panic: close of closed channel
	// 不能向一个已经关闭的通道发送消息                                 // panic: send on closed channel
	// 向一个已经关闭的通道发送数据：intChannel <- 1                    // panic: send on closed channel
	println(<-intChannel)

	// 关闭通道，一定是关闭一个可写的通道（cannot close receive-only channel），编译不通过

	// 无法判断一个只写通道是否关闭

	// 可以向一个已经关闭的带有缓冲区的通道正常读取缓冲区里的数据
}
