package basic

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	theFunc := func() {
		println("1 s...")
	}

	// 方式一
	RunAtEverySecond(theFunc)
	// 方式二
	// RunAtEverySecond2(theFunc)
	// 方式三：非阻塞式的单次调用
	// time.AfterFunc(1 * time.Second, theFunc)
}

func RunAtEverySecond(theFunc func()) {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		theFunc()
	}
}

func RunAtEverySecond2(theFunc func()) {
	for {
		select {
		case <-time.After(1 * time.Second):
			theFunc()
		}
	}
}
