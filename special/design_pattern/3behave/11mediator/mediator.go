package _1mediator

import (
	"fmt"
	"reflect"
)

// 有一个较为复杂的对话框，对话框有多个 tab，每个 tab 都对应不同的逻辑的相关组件
// 参考给出的样例如下；感觉大多数设计模式就好像是把某个领域，特定业务场景的最佳实践归为的一种设计模式

// Input 输入框
type Input string

func (i Input) String() string {
	return string(i)
}

// Select 选择框
type Select string

func (i Select) Selected() string {
	return string(i)
}

// Button 按钮
type Button struct {
	onClick func()
}

func (b *Button) SetOnClick(onClick func()) {
	b.onClick = onClick
}

// Dialog 对话框
type Dialog struct {
	Tab                 *Select
	LoginButton         *Button
	RegButton           *Button
	UsernameInput       *Input
	PasswordInput       *Input
	RepeatPasswordInput *Input
}

func (d *Dialog) HandleEvent(component interface{}) {
	switch {
	case reflect.DeepEqual(component, d.Tab):
		if d.Tab.Selected() == "登录" {
			fmt.Println("select login")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
		} else if d.Tab.Selected() == "注册" {
			fmt.Println("select register")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			fmt.Printf("show: %s\n", d.RepeatPasswordInput)
		}
	}
}
