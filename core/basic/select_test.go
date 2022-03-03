package basic

import (
	"testing"
	"time"
)

// select 语法定义和通道一起使用，常运行在无条件的循环中
// 语法：Select case must be receive, send or assign receive
//
// 原理：执行 select 逻辑，会遍历所有 case，在所有符合条件（准备接收元素通道 的 灌输表达式；准备吐出元素通道 的 读取表达式）的 case 中随机选择一个执行；
//   如果没有符合条件的，就会走 default 逻辑（如果定了的话）
//
// 注意，遍历 case，意味着会尝试执行 case 对应的通道表达式（不是实际执行）；这里想强调的意思是说
// - 如果是一个通道输值表达式 且 通道为 nil，那么就会引发 panic
// - 如果是一个 “双重” 通道的灌输数值表达式，可能存在很大的问题

func TestSelect(t *testing.T) {
	// basicTest(t)

	// doubleChanTest(t)

	// multiRecvTest(t)
	multiRecvBetterTest(t)
}

// 非常好的，能够说明 select 基本原理和用法的例子
func basicTest(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func() { c1 <- 1 }()
	go func() { c2 <- 2 }()
	go func() { c3 <- 3 }()

	// 1、如果去掉上面的 Goroutine，那么下边每次执行的就是 default
	// 2、select 会阻塞至任意一个 case 满足条件
	// 3、case 的选择是随机的，且会因为其他 Goroutine 和主 Goroutine 的执行顺序的不定性，会有 default 的执行
	// 如果在 select 上添加 time.Sleep(xxx)，能够确保下边每次执行 select 中的 case，而不是 default
	select {
	case v := <-c1:
		t.Logf("从通道1读取值：%d\n", v)
	case v := <-c2:
		t.Logf("从通道2读取值：%d\n", v)
	case v := <-c3:
		t.Logf("从通道3读取值：%d\n", v)
	default:
		t.Log("No match and default run")
	}
}

func doubleChanTest(t *testing.T) {
	c := make(chan struct{}, 1)
	m := make(chan struct{}, 0)

	go func() { m <- struct{}{} }()

	c <- <-m
	t.Logf("hei")
}

// 这个例子不能清楚的说明问题，下面会做改进，但是就这个例子你可能会认为 1 和 2 随机打印，打印到第 5 次后退出
// 实际是，一定会以死锁的结果收尾
// 这里主要是突出 case 的评估，为了监测 case 的直接通道操作是否可行，一定会先获取向通道传递的值
// 也就是最终卡死是卡死在了这个获取值的操作上
func multiRecvTest(t *testing.T) {
	c := make(chan struct{}, 5)
	pre1 := make(chan struct{}, 5)
	pre2 := make(chan struct{}, 5)

	for i := 0; i < 5; i++ {
		pre1 <- struct{}{}
		pre2 <- struct{}{}
	}

	for {
		select {
		case c <- <-pre1:
			t.Log(1)
		case <-pre2:
			t.Log(2)
		default:
			return
		}
	}
}

// 所有 case 的通道估值方法每次都会执行
// 和上面的样例结合起来重点说明，就是通道间传值表达式，给值的那个通道给出的值可能会因为条件成立，
// 但是因为 case 的随机性而没有被执行，就会导致给值通道的给出的值丢失
func multiRecvBetterTest(t *testing.T) {
	var c1 = make(chan int, 1)
	var c2 = make(chan int, 1)
	i1 := 0
	i2 := 0

	v1 := func() int {
		i1++
		t.Log("v1 send", i1)
		return i1
	}
	v2 := func() int {
		i2++
		t.Log("v2 send", i2)
		return i2
	}

	for range time.Tick(time.Second) {
		t.Log("run...")
		select {
		case c1 <- v1():
			t.Log("c1 receive", <-c1)
		case c2 <- v2():
			t.Log("c2 receive", <-c2)
		}
	}
}
