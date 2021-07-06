package help

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

/*
使用 go 的一个三方库来解析 cron 表达式任务

GitHub：github.com/robfig/cron/v3
Doc：https://pkg.go.dev/github.com/robfig/cron/v3

// 回顾原来的笔记
// Linux 中使用 cron 和 java 中有不同
// 特殊符号，并不是所有类库都实现了，这里的 go 三方库，支持 ?
// 查看文档的方式记录

和 cron 表达式规范有出入的地方在于，秒是可选的
*/
type CronJob []cronJob

type cronJob struct {
	Name string       // 对要执行任务的中文描述
	Cron string       // cron 表达式
	Job  func() error // 要执行的任务（可以直接使用 cron.Job 类型，但是希望将单独的日志提示提出来）
}

// 定义任务要执行的任务 - 需要实现的接口 cron.Job
func (js cronJob) Run() {
	if err := js.Job(); err != nil {
		fmt.Printf("error occur! when execute cron job：【%s】【%s】【%s】\n",js.Cron, js.Name, err)
	} else {
		fmt.Printf("execute cron job：【%s】【%s】\n",js.Cron, js.Name)
	}
}

// 开始定时任务
func (js CronJob) Start() (err error) {
	var c = cron.New()
	defer func() {
		if err == nil {
			c.Start()
		}
	}()

	for _, j := range js {
		if _, err = c.AddJob(j.Cron, j); err != nil {
			fmt.Printf("error occur! when add a new cron job 【cron：%s】【name：%v】【err：%s】\n",j.Cron, j.Name, err)
			return
		}
		fmt.Printf("add a new cron job：【%s】【%s】\n",j.Cron, j.Name)
	}
	return
}
