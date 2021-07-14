package pool

import "testing"

/*
这里先做一个引出，就以处理 http 请求的 server 为背景，就是说在一个可能有无限并发可能的逻辑处，需要对 Goroutine 的数量进行控制
实例来源于 《Effective Go》
*/

func TestPre(t *testing.T) {
	failOne()
	failTwo()
	successOne()
	successTwo()
	successThree()
}

// gg：能够限制 cpu 同一时间执行的最大 Goroutine 数量，但是没有限制最大的 Goroutine 数量
func failOne() {
	// var sem = make(chan int, MaxOutstanding)
	// func handle(r *Request) {
	// 	sem <- 1 // 等待活动队列清空。
	// 	process(r) // 可能需要很长时间。
	// 	<-sem // 完成；使下一个请求可以运行。
	// }
	// func Serve(queue chan *Request) {
	// 	for {
	// 		req := <-queue
	// 		go handle(req) // 无需等待 handle 结束。
	// 	}
	// }
}

// gg：因闭包特性导致并发安全问题，req 变量是多个 Goroutine 共享的
func failTwo() {
	// func Serve(queue chan *Request) {
	// 	for req := range queue {
	// 		sem <- 1
	// 		go func() {
	// 			process(req) // Buggy; see explanation below.
	// 			<-sem
	// 		}()
	// 	}
	// }
}

func successOne() {
	// func Serve(queue chan *Request) {
	// 	for req := range queue {
	// 		sem <- 1
	// 		go func(req *Request) {
	// 			process(req)
	// 			<-sem
	// 		}(req)
	// 	}
	// }
}

func successTwo() {
	// func Serve(queue chan *Request) {
	// 	for req := range queue {
	// 		req := req // 为该 Go 程创建 req 的新实例。
	// 		sem <- 1
	// 		go func() {
	// 			process(req)
	// 			<-sem
	// 		}()
	// 	}
	// }
}

func successThree() {
	// func handle(queue chan *Request) {
	// 	for r := range queue {
	// 		process(r)
	// 	}
	// }
	// func Serve(clientRequests chan *Request, quit chan bool) {
	// 	// Start handlers
	// 	for i := 0; i < MaxOutstanding; i++ {
	// 		go handle(clientRequests)
	// 	}
	// 	<-quit // Wait to be told to exit.
	// }
}
