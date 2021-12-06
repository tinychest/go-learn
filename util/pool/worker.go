package pool

import "fmt"

// 工人，执行任务的基础单位
type Worker struct {
	// 接收任务
	taskChan <-chan Task
	// 内部发生错误，向外部告知错误信息（整个逻辑将会结束）
	errChan chan<- error
	// 输出结果
	resultChan chan<- interface{}
	// 外部通知结束的通道
	closeSignal <-chan struct{}
}

func NewWorker(input <-chan Task, err chan<- error, output chan<- interface{}, close <-chan struct{}) *Worker {
	return &Worker{
		taskChan:    input,
		errChan:     err,
		resultChan:  output,
		closeSignal: close,
	}
}

func (w *Worker) Start() {
	go w.start()
}

func (w *Worker) start() {
	var (
		task   Task
		result interface{}
		err    error
	)

	// 假如执行任务发生 panic，可能会导致核心机制的结果接收 goroutine 无法结束
	defer w.panicHandle()

	for task = range w.taskChan {
		// 退出：内部发生错误
		if result, err = task(); err != nil {
			w.errChan <- err
			return
		}
		// 退出：外部要求
		select {
		case <-w.closeSignal:
			return
		default:
		}
		// 退出：不符合机制要求，任务执行的结果不能为空
		if result == nil {
			w.errChan <- fmt.Errorf("任务执行结果为空")
			return
		}
		w.resultChan <- result
	}
}

func (w *Worker) panicHandle() {
	if errObj := recover(); errObj != nil {
		w.errChan <- errObj.(error)
	}
}
