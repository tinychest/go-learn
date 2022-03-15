package concurrency

import (
	"fmt"
	"testing"
)

var totalTaskSum = 4
var currentCompleteSum = 0

func TestGoroutine(t *testing.T) {
	channel := make(chan int)
	go commonLogic(channel)
	go commonLogic(channel)
	go commonLogic(channel)
	go commonLogic(channel)
}

func commonLogic(channel chan int) {
	// 确保通道关闭的逻辑（要使用 defer 修饰，防止执行逻辑的时候，发生异常）
	defer func() {
		currentCompleteSum++
		if currentCompleteSum == totalTaskSum {
			fmt.Printf("第 %d 执行完任务，关闭通道\n", currentCompleteSum)
			close(channel)
		}
	}()

	// 执行逻辑
	println("执行任务...")
}

// Go 语法注意点：后边不能直接接要执行的语句，一定得放到 函数 中
func goroutineBasicTest(t *testing.T) {
	// 编译通过
	// go t.Log("123")
	// 编译不通过？上面 go 后边跟的是函数调用，下面不是
	// go intChannel <- 1
}
