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

	// resultChan、errChan 自然回收

	defer c.cancel()

	// 接收结果
	go func() {
		defer close(output)
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

	// 发放任务
	go func() {
		// NOTE 一定要在这里关闭，实际出现过下边向 taskChan 发送数据的时候，taskChan 关闭了（几率非常小）
		defer close(taskChan)
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
