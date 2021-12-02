package _command

import "fmt"

/* 示例1：要执行的函数看作一个命令，函数封装成实例，放入队列，延迟执行 */

type ICommand interface {
	Execute() error
}

type StartCommand struct {}
type EndCommand struct{}

func NewStartCommand(/* 正常情况下，这里会有一些参数 */) *StartCommand {
	return &StartCommand{}
}

func NewArchiveCommand(/* 正常情况下，这里会有一些参数 */) *EndCommand {
	return &EndCommand{}
}

func (c *StartCommand) Execute() error {
	fmt.Println("start...")
	return nil
}

func (c *EndCommand) Execute() error {
	fmt.Println("end...")
	return nil
}

/* 示例2：类比示例1，这里直接使用函数，不用结构体实例（对象） */

type ICommand2 func() error

func StartCommand2() error {
	fmt.Println("start...")
	return nil
}

func EndCommand2() error {
	fmt.Println("end...")
	return nil
}
