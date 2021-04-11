package concurrency

import (
	"errors"
	"fmt"
)

// 借助辅助双通道，实现 channel_bgtask_test.go 中描述的业务逻辑
// 仅仅多借助了一个辅助通道，程序实现复杂度大幅降低，并且是完美符合业务需求

type TaskExecModel int

var (
	Sync  TaskExecModel = 0 // 提交任务时，假如通道已满则等待
	Async TaskExecModel = 1 // 提交任务时，假如通道已满则返回错误信息
)

type bgTask struct {
	execModel TaskExecModel

	taskChannel chan interface{}
	taskHandler func(interface{})

	closeChannel         chan int
	closeCompleteChannel chan int
}

type IBgTask interface {
	// 启动
	Start()

	// 提交任务
	SubmitTask(task interface{}) error

	// 关闭（阻塞 - 如果任务还没有执行完），如果 web 框架使用的是 beego 那么一定要在配置文件中加上 Graceful = true
	Close()
}

func NewBgTask(taskChannel chan interface{}, taskHandler func(interface{}), excelModel TaskExecModel) IBgTask {
	return &bgTask{
		execModel:            excelModel,
		taskChannel:          taskChannel,
		taskHandler:          taskHandler,
		closeChannel:         make(chan int, 0),
		closeCompleteChannel: make(chan int, 0),
	}
}

func (bg *bgTask) SubmitTask(task interface{}) error {
	switch bg.execModel {
	case Sync:
		bg.taskChannel <- task
		return nil
	case Async:
		select {
		case bg.taskChannel <- task:
			return nil
		default:
			return errors.New(fmt.Sprintf("同一时间最多允许存在 %d 个表格导出任务，当前任务队列已满", cap(bg.taskChannel)+1))
		}
	default:
		// never reach here
		return nil
	}
}

func (bg *bgTask) Start() {
	go bg.start()
}

func (bg *bgTask) Close() {
	bg.close()
}

func (bg *bgTask) start() {
loop:
	for {
		select {
		case task := <-bg.taskChannel:
			bg.taskHandler(task)
		case <-bg.closeChannel:
			break loop
		}
	}
	for len(bg.taskChannel) != 0 {
		bg.taskHandler(<-bg.taskChannel)
	}
	<-bg.closeCompleteChannel
}

func (bg *bgTask) close() {
	bg.closeChannel <- 1
	bg.closeCompleteChannel <- 1
}
