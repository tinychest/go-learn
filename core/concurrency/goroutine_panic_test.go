package concurrency

import (
	"testing"
	"time"
)

// 子 Goroutine panic 父 Goroutine 直接终止
func TestGoRoutinePanic(t *testing.T) {
	go func() {
		t.Fatal("子 Goroutine 慌了")
	}()

	println("主 Goroutine 开始睡")
	time.Sleep(2 * time.Second)
	println("主 Goroutine 睡醒了")
}
