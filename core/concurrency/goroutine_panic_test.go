package concurrency

import (
	"testing"
	"time"
)

// 子协程 panic 父协程 直接终止
func TestGoRoutinePanic(t *testing.T) {
	go func() {
		t.Fatal("子协程慌了")
	}()

	println("主协程开始睡")
	time.Sleep(2 * time.Second)
	println("主协程睡醒了")
}
