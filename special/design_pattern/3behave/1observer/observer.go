package _observer

import "fmt"

// 最简单、基础的 样例

// ISubject 发布者
type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

// IObserver 观察者
type IObserver interface {
	Receive(msg string)
}

// Subject 发布实现
type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Remove(observer IObserver) {
	for i, v := range s.observers {
		if v == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, v := range s.observers {
		v.Receive(msg)
	}
}

// Observer1 观察者实现1
type Observer1 struct{}

func (o *Observer1) Receive(msg string) {
	fmt.Println("观察者01 接收到消息" + msg)
}

// Observer2 观察者实现2
type Observer2 struct{}

func (o *Observer2) Receive(msg string) {
	fmt.Println("观察者02 接收到消息" + msg)
}
