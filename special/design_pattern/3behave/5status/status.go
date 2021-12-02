package _status

import "fmt"

/*
以审批流程的业务场景来说，每个环节就是不同的状态，而审批动作（批准、驳回）就是不同的事件
*/

// Machine 状态机
type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetState() string {
	return m.state.Name()
}

func (m *Machine) Approve() {
	m.state.Approve(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// IState 状态流转
type IState interface {
	Reject(m *Machine)
	Approve(m *Machine)
	Name() string
}

// 状态定义

type Step01 struct{}
type Step02 struct{}
type FinalStep struct{}

func (s *Step01) Reject(_ *Machine) {
	panic("no mean for reject first step")
}

func (s *Step01) Approve(m *Machine) {
	m.SetState(new(Step02))
}

func (s *Step01) Name() string {
	return "STEP01"
}

func (s *Step02) Reject(m *Machine) {
	// 相应的业务逻辑
}

func (s *Step02) Approve(m *Machine) {
	fmt.Println("审批通过")
	m.SetState(new(FinalStep))
}

func (s *Step02) Name() string {
	return "STEP02"
}

func (s *FinalStep) Reject(m *Machine) {
	// 相应的业务逻辑
}

func (s *FinalStep) Approve(m *Machine) {
	panic("no mean for approve final step")
}

func (s *FinalStep) Name() string {
	return "FINAL"
}