package concurrency

import (
	"testing"
)

// 描述一个业务场景：
// 有一个 Goroutine 在后台监听任务，监听 任务通道 和 结束信号通道
// 任务通道 和 结束信号 通道都有数据时以任务通道为主，即任务通道没有数据，结束信号通道有数据才结束

// 背景条件1：假如收到了关闭信号，那么就不会再产生新任务了（任务池还有的，不影响，依旧存在）
// 背景条件2：关闭信号只会发送一次
func TestChannelOnlineBgTaskTest(t *testing.T) {
	var taskChannel chan interface{}
	var closeChannel chan int

	version1(taskChannel, closeChannel)
}

// 没有体现出执行任务的优先级要大于结束 Goroutine
func version1(taskChannel chan interface{}, closeChannel chan int) {
	for {
		select {
		case task := <-taskChannel:
			println(task)
		case <-closeChannel:
			break
		}
	}
}

// 乍以为实现了，但是实际是如果当前没有任务就在只等待关闭信息，即不接收新任务了
func version2(taskChannel chan interface{}, closeChannel chan int) {
	for {
		select {
		case task := <-taskChannel:
			println(task)
		default:
			select {
			case <-closeChannel:
				break
			}
		}
	}
}

// 仔细想想，题目是存在悖论的，你希望有一个优先级（工作状态时），你又希望来哪个就执行哪个（空闲状态时）
// 解释：如果你希望来哪个执行哪个那么门面肯定是 case case，但是你又希望任务优先所以是 case default
// 只从 case case 的情况思考，如何让任务的优先级优先
// 结果：乍一看，已经实现了需求，反例：
// 当前任务池有 3 个任务，执行完第一个发现来了结束信号，正好也判断了结束信号，但是发现还有任务，于是继续执行任务了...没错，结束信号被丢弃了
func version3(taskChannel chan interface{}, closeChannel chan int) {
	for {
		select {
		case task := <-taskChannel:
			println(task)
		case <-closeChannel:
			select {
			case task := <-taskChannel:
				println(task)
			default:
				return
			}
		}
	}
}

// 解决结束信号丢失 - 额，补上呗；
// 但是在子 Goroutine 收到关闭信号的同时，主 Goroutine 等待结束信号接收了，主 Goroutine 就结束了，子 Goroutine 全部关闭
// 借助辅助变量来实现：稍微构思一下，就会发现，实现起来太复杂了，不现实
func version4(taskChannel chan interface{}, closeChannel chan int) {
	for {
		select {
		case task := <-taskChannel:
			println(task)
		case <-closeChannel:
			select {
			case task := <-taskChannel:
				closeChannel <- 1 // 续上结束信号
				println(task)
			default:
				return
			}
		}
	}
}

// 从 case default 的门面来思考
func finalVersion1(taskChannel chan interface{}, closeChannel chan int) {
	for {
		select {
		// 有任务执行任务，其他什么都不管
		case task := <-taskChannel:
			println(task)
		default:
			select {
			// 没任务了，看下要关闭么
			case <-closeChannel:
				return
			// 空闲状态，平等的接收任务或者关闭信号
			default:
				select {
				case task := <-taskChannel:
					println(task)
				case <-closeChannel:
					return
				}
			}
		}
	}
}

// 你会发现上面的明显可以优化
// 你还可以发现，只从通道的最终等待状态来分析，也是可以推出下面这样的解的
func finalVersion2(taskChannel chan interface{}, closeChannel chan int) {
	for {
		select {
		// 有任务执行任务，其他什么都不管
		case task := <-taskChannel:
			println(task)
		// 没任务了，就等任务，或者关闭
		default:
			select {
			case task := <-taskChannel:
				println(task)
			case <-closeChannel:
				return
			}
		}
	}
}

// 总结：finalVersion2 是最简单实用的写法，但是仔细想一想，假如是空闲状态下，任务和关闭信号同时来了，那么该种写法并不能保证同时，是任务先执行
// 当然，怎么可能存在“同时”，只是理论上分析而已。推荐采用 channel_bgtask2_test.go 中的封装模块，已在项目中实战
