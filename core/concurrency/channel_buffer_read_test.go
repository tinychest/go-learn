package concurrency

import (
	"testing"
)

// 不带缓存的通道
// 读取就是通道的基础语法，这里不做描述

// 带有缓存的通道：make(chan int, n)
// cap 代表缓冲区大小（同数组，没有扩容一说），len 代表缓冲区当前元素的个数（同切片，元素是添加和删除的）
// 通道缓冲区大小默认是 0，即没有缓冲区：make(chan int)、make(chan int, 0)
func TestChannelRead(t *testing.T) {
	intBufferChannel := make(chan int, 4)

	// 方式一：for select break
	// 结束时机：break's condition

	// 这里边有个坑，就是说直接调用 break 等于没调用，以为 break 相当于跳出 select，需要指定层级跳出
	// select：选择 case 中任意一个达成条件的执行，没有达成条件的就一直等待
	// for select：只要 case 中任意一个达成条件就会执行，没有达成条件的就一直等待
forLoop:
	for {
		select {
		case intValue := <-intBufferChannel:
			t.Logf("从通道读取：%d\n", intValue)
		default:
			if len(intBufferChannel) == 0 {
				break forLoop
			}
		}
	}

	// 方式二：for select（推荐）
	// 结束时机：for condition
	estimateCircleCount := 3
	for index := 0; index < estimateCircleCount; index++ {
		select {
		case intValue := <-intBufferChannel:
			t.Logf("从通道读取：%d\n", intValue)
		}
	}

	// 方式三：for range（推荐）
	// 结束时机：chan close

	// 这里和 for range slice 不同的是，for range chan 不能定义两个返回值的接收形式
	for intValue := range intBufferChannel {
		t.Logf("从通道读取：%d\n", intValue)
	}
}
