package _observer

import (
	"fmt"
	"testing"
	"time"
)

func subscribFunc1(param1 string, param2 int) {
	fmt.Println(param1, param2)
}

func subscribFunc2(param1 string, param2 int) {
	fmt.Println(param1, param2)
}

func TestBus(t *testing.T) {
	bus := NewAsyncEventBus()

	_ = bus.Subscribe("topic-01", subscribFunc1)
	_ = bus.Subscribe("topic-02", subscribFunc2)

	bus.Publish("topic-01", "a", 1)
	bus.Publish("topic-02", "b", 2)

	time.Sleep(time.Second)
}
