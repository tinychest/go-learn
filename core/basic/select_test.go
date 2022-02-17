package basic

import (
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	// for {
	//     basicTest(t)
	// }

	// caseChannelReceiverTest(t)
}

// 下面的代码逻辑，乍看没有问题，实际会被阻塞住
// 那是因为没有好好理解 select case 真正的原理，select case 会在所有符合条件的里边随机取一个
// 这边存在，我怎么知道是否符合条件，那么就只能去执行判断一下，然而，当执行 <-getSignal 的时候就卡住了
func TestPit(t *testing.T) {
	taskSignal := make(chan struct{})
	getSignal := make(chan struct{})
	outSignal := make(chan struct{})

	go func() {
		for task := range taskSignal {
			t.Log("收到任务，执行并输出结果")
			getSignal <- task
		}
	}()

	select {
	case taskSignal <- struct{}{}:
		t.Log("发送任务")
	case outSignal <- <-getSignal:
		t.Log("传递结果")
	}
}

// Select 语法注意点：Select case must be receive, send or assign receive
// case 语法注意点：子句必须是一个通信操作，要不是发送，要不是接收
// default 语法注意点：如果所有 case 子句都不符合执行条件，那就会执行 default 子句
// Go 对 Goroutine 的状态有一个监听模式，如果所有携程都阻塞了，则程序会直接终止：fatal error: all goroutines are asleep - deadlock!
func basicTest(t *testing.T) {
	intChannel1 := make(chan int)
	intChannel2 := make(chan int)
	intChannel3 := make(chan int)

	go func() {
		intValue := 1
		// t.Logf("向通道 1 发送值：%d\n", intValue)
		intChannel1 <- intValue
	}()
	go func() {
		intValue := 1
		// t.Logf("向通道 2 发送值：%d\n", intValue)
		intChannel2 <- intValue
	}()
	go func() {
		intValue := 1
		// t.Logf("向通道 3 发送值：%d\n", intValue)
		intChannel3 <- intValue
	}()

	// 1、如果去掉上面的 Goroutine，那么下边每次执行的就是 default
	// 2、select 会阻塞至任意一个 case 满足条件
	// 3、case 的选择是随机的，且会因为其他 Goroutine 和主 Goroutine 的执行顺序的不定性，会有 default 的执行
	// 如果在 select 上添加 time.Sleep(xxx)，能够确保下边每次执行 select 中的 case，而不是 default
	select {
	case receiveInt := <-intChannel1:
		t.Logf("从通道1读取值：%d\n", receiveInt)
	case receiveInt := <-intChannel2:
		t.Logf("从通道2读取值：%d\n", receiveInt)
	case receiveInt := <-intChannel3:
		t.Logf("从通道3读取值：%d\n", receiveInt)
	default:
		t.Log("No match and default run")
	}
}

// 一个发送者，多个 select 的接收者（其实调过来也完全没差，都能体现关键的语法）
func caseChannelReceiverTest(t *testing.T) {
	intChannel := make(chan int)

	// 和上面的不同，这里定义接收
	go func() {
		time.Sleep(10 * time.Second)
		intValue := <-intChannel
		t.Logf("从通道接收：%d\n", intValue)
	}()

	// 在接收者中定义阻塞，select 是会阻塞住的
	// 这里可以体现出 case 选择的随机性
	// 这里可以体现出 并发 的随机性：发送这条语句的打印，在接收语句打印之后
	select {
	case intChannel <- 1:
		t.Logf("向通道发送：%d\n", 1)
	case intChannel <- 2:
		t.Logf("向通道发送：%d\n", 2)
	}
}
