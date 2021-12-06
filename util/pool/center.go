package pool

import "context"

type TaskCenter interface {
	// Start 以预定数量可复用的 goroutine，同时执行任务
	Start(func(result interface{}))
	// Close 外部希望结束整个任务（例如，设定超时时间）
	Close()
	// Error 获取机制执行过程中发生的错误
	Error() error
}

type taskCenter struct {
	inner
}

func NewTaskCenter(total int, getter func(index int) Task) TaskCenter {
	ctx, cancel := context.WithCancel(context.Background())

	return &taskCenter{
		inner: inner{
			ctx:    ctx,
			cancel: cancel,
			tasker: newTasker(total, getter),
		},
	}
}

func (c *taskCenter) Start(collector func(interface{})) {
	output := make(chan interface{})

	go start(c.inner, output)
	for v := range output {
		collector(v)
	}
}

func (c *taskCenter) Close() {
	c.cancel()
}

func (c *taskCenter) Error() error {
	if v := c.ctx.Value(ErrKey); v != nil {
		return v.(error)
	}
	return nil
}
