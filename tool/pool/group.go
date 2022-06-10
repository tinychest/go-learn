package pool

// Group 管理工人组，本身不和通道的数据交互
type Group struct {
	// 告知 Worker 结束的信号通道
	closeSignal chan struct{}
	// workers 工作组的工人
	workers []*Worker
}

func NewGroup(input <-chan Task, output chan<- *Result, sum int) (g *Group) {
	g = &Group{ closeSignal: make(chan struct{}) }

	ws := make([]*Worker, sum)
	for i := 0; i < sum; i++ {
		ws[i] = NewWorker(input, output, g.closeSignal)
	}
	g.workers = ws

	return
}

func (g *Group) Start() {
	for _, v := range g.workers {
		v.Start()
	}
}

func (g *Group) Close() {
	close(g.closeSignal)
}
