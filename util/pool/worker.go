package pool

// Worker 工人，每个都会对应一个 Goroutine，执行任务的基础单位
type Worker struct {
	// 接收任务
	taskChan <-chan Task
	// 输出结果
	resultChan chan<- *Result
	// 外部（工作组）通知结束的通道
	closeSignal <-chan struct{}
}

func NewWorker(taskChan <-chan Task, resultChan chan<- *Result, closeSignal <-chan struct{}) *Worker {
	return &Worker{
		taskChan:    taskChan,
		resultChan:  resultChan,
		closeSignal: closeSignal,
	}
}

func (w *Worker) Start() {
	go w.start()
}

func (w *Worker) start() {
	defer w.panicHandle()

	for task := range w.taskChan {
		res, err := task()
		select {
		case <-w.closeSignal:
			return
		default:
		}

		select {
		case <-w.closeSignal:
			return
		case w.resultChan <- &Result{res: res, err: err}:
		}
	}

	// TODO
	// fmt.Println("Worker exit...")
}

func (w *Worker) panicHandle() {
	if errObj := recover(); errObj != nil {
		select {
		case <-w.closeSignal:
			return
		case w.resultChan <- &Result{err: errObj.(error)}:
		}
	}
}
