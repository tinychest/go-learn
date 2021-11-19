package _observer

import (
	"fmt"
	"reflect"
	"sync"
)

// Bus 总线
type Bus interface {
	// Subscribe 订阅 handle 处理函数
	Subscribe(topic string, handler interface{}) error
	// Publish 发布 args 是处理函数的参数
	Publish(topic string, args ...interface{})
}

// AsyncEventBus 总线实现
type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := bus.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handlers[topic] = handler

	return nil
}

func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}
