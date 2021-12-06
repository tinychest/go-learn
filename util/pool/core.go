package pool

import (
	"context"
	"runtime"
)

var maxGoroutineSum = runtime.GOMAXPROCS(8)

const ErrKey = "ERROR"

type inner struct {
	// context 机制（存储任务执行过程中的错误信息）
	ctx context.Context
	// 机制关闭的核心
	cancel context.CancelFunc
	// 任务池
	tasker Tasker
}

// 核心方法
func start(c inner, output chan<- interface{}) {
	var (
		taskChan   = make(chan Task)
		errChan    = make(chan error)
		resultChan = make(chan interface{})
		group      = NewWorkGroup(taskChan, errChan, resultChan, maxGoroutineSum)
	)
	group.Start()
	defer group.Close()
	defer close(taskChan)
	defer close(resultChan)
	defer close(output)

	defer c.cancel()

	// 接收结果
	go func() {
		for i := 0; i < c.tasker.GetTotal(); i++ {
			select {
			case <-c.ctx.Done():
				return
			default:
				result := <-resultChan
				if result == nil {
					return
				}
				output <- result
			}
		}
		c.cancel()
	}()

	// 发送任务
	go func() {
		for i := 0; i < c.tasker.GetTotal(); i++ {
			select {
			case <-c.ctx.Done():
				return
			default:
				select {
				case taskChan <- c.tasker.GetTask(i):
				case <-c.ctx.Done():
					return
				}
			}
		}
	}()

	// 控制中心
	for {
		select {
		// 任务执行完毕、外部要求结束
		case <-c.ctx.Done():
			return

		// 内部任务执行出错
		case err := <-errChan:
			c.ctx = context.WithValue(c.ctx, ErrKey, err)
			return
		}
	}
}
