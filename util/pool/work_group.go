package pool

// 作为所有工人的启动和关闭入口，用于管理工人组，本身不和通道的数据交互【只含有自己管理的通道】
type WorkerGroup struct {
	// 告知 Worker 结束的信号通道
	closeSignal chan struct{}
	// workers 工作组的工人（一个工人对应一个 goroutine）
	workers []*Worker
}

// err 输出 Worker 执行任务出错信息的通道
func NewWorkGroup(input <-chan Task, err chan<- error, output chan<- interface{}, sum int) *WorkerGroup {
	g := &WorkerGroup{
		closeSignal: make(chan struct{}),
		workers:     nil,
	}

	ws := make([]*Worker, sum)
	for i := 0; i < sum; i++ {
		ws[i] = NewWorker(input, err, output, g.closeSignal)
	}
	g.workers = ws

	return g
}

func (g *WorkerGroup) Start() {
	for _, v := range g.workers {
		v.Start()
	}
}

func (g *WorkerGroup) Close() {
	close(g.closeSignal)
}
