package help

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

// 定时任务
type cronJob struct {
	Name string
	Cron string // 使用`robfig/cron`解析时，默认不包括秒
	Job  func() error
}

type CronJobs []cronJob

// 开始定时任务
func (js CronJobs) Start() error {
	var c = cron.New()
	for index, j := range js {
		fmt.Printf("添加第[%d]个定时[%s]任务[%s]\n", index+1, j.Cron, j.Name)
		if _, err := c.AddJob(j.Cron, j); err != nil {
			return err
		}
	}
	c.Start()
	return nil
}

// 实现接口 cron.Job
func (js cronJob) Run() {
	fmt.Printf("开始执行定时[%s]任务[%s]：\n", js.Name, js.Cron)

	if err := js.Job(); err != nil {
		fmt.Printf("执行出错-定时[%s]任务[%s]：%s\n", js.Name, js.Cron, err)
	} else {
		fmt.Printf("执行成功-定时[%s]任务[%s]\n", js.Name, js.Cron)
	}
}
