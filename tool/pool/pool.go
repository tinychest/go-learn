package pool

import (
	"context"
	"errors"
	"runtime"
	"time"
)

var maxGoroutineSum = runtime.GOMAXPROCS(8)

const ErrKey = "ERROR"

type Center interface {
	// Start 以预定数量可复用的 goroutine，同时执行任务
	Start(callback func(result interface{}), timeout time.Duration)
	// Close 外部希望结束整个任务（例如，设定超时时间）
	Close()
	// Error 获取机制执行过程中发生的错误
	Error() error
}

type center struct {
	// 状态
	state int
	// 存储任务执行过程中的错误信息
	ctx context.Context
	// 关闭按钮（任务执行完毕、任务执行出错）
	cancel context.CancelFunc
	// 任务池
	tasker Tasker
}

func NewCenter(total int, getter func(index int) Task) Center {
	return &center{
		tasker: newTasker(total, getter),
	}
}

func (c *center) Start(callback func(result interface{}), timeout time.Duration) {
	if timeout == 0 {
		c.ctx, c.cancel = context.WithCancel(context.Background())
	} else {
		c.ctx, c.cancel = context.WithTimeout(context.Background(), timeout)
	}
	defer c.cancel()

	var (
		output = make(chan interface{})
		sum    = 0
	)
	go start(c, output)
	for v := range output {
		sum++
		callback(v)
	}

	// 任务超时（没有发生任何错误，却没有收集到足够的结果）
	if c.Error() == nil && sum != c.tasker.Total() {
		c.ctx = context.WithValue(c.ctx, ErrKey, errors.New("center timeout"))
	}
}

func start(c *center, output chan<- interface{}) {
	var (
		taskChan   = make(chan Task)
		resultChan = make(chan *Result)
		group      = NewGroup(taskChan, resultChan, maxGoroutineSum)
	)
	group.Start()
	defer group.Close()

	// 不关闭 resultChan 没事，因为结束时，肯定没有接收者了，那么发送者肯定无法发送成功，所以只能走退出了

	// 接收结果
	go func() {
		defer close(output)
		defer c.cancel()
		for i := 0; i < c.tasker.Total(); i++ {
			select {
			case <-c.ctx.Done():
				return
			default:
			}

			select {
			case <-c.ctx.Done():
				return
			case result := <-resultChan:
				if result.err != nil {
					c.ctx = context.WithValue(c.ctx, ErrKey, result.err)
					return
				}
				if result.res == nil {
					c.ctx = context.WithValue(c.ctx, ErrKey, errors.New("任务执行结果为空"))
					return
				}
				output <- result.res
			}
		}
	}()

	// 发放任务
	go func() {
		// NOTE 一定要在这里关闭，出现过下边向 taskChan 发送数据的时候，taskChan 关闭了
		defer close(taskChan)
		for i := 0; i < c.tasker.Total(); i++ {
			select {
			case <-c.ctx.Done():
				return
			default:
			}

			select {
			case <-c.ctx.Done():
				return
			case taskChan <- c.tasker.Task(i):
			}
		}
	}()

	<-c.ctx.Done()

	// TODO
	// fmt.Println("Center exit...")
}

func (c *center) Close() {
	c.cancel()
}

func (c *center) Error() error {
	if v := c.ctx.Value(ErrKey); v != nil {
		return v.(error)
	}
	return nil
}
