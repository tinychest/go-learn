package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"testing"
)

// TODO ?
func TestDeffer(t *testing.T) {
	var i = 1

	// 最后执行：1（立即故值，延迟执行）
	defer fmt.Println("zero", i)
	// 后执行：2
	defer func() {fmt.Println("one", i)}()
	// 先执行：1（立即故值，延迟执行）
	defer func(param int) {fmt.Println("two", param)}(i)

	i = 2
}

// 测试 cancel 执行后，ctx.Done 可以被多个监听者响应么
func TestM(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	select {
	case <- ctx.Done():
		fmt.Println("ok")
	}

	select {
	case <- ctx.Done():
		fmt.Println("ok")
	}
}

func TestJsonMarMar(t *testing.T) {
	j := `{"name":"xiaoming"}`

	var r []byte
	var err error
	var i2 interface{}

	// +
	if r, err = json.Marshal(j); err != nil {
		panic(err)
	} else {
		fmt.Println(string(r))
	}

	// json 串经过 Marshal 无法再 Unmarshal 了
	if err = json.Unmarshal(r, i2); err != nil {
		panic(err)
	} else {
		fmt.Println(i2)
	}
}

// defer 定义的位置，并不是去闭包的死值，是动态的
func TestDefer(t *testing.T) {
	var i = 1
	var err = errors.New("123")
	defer func() {
		fmt.Println(i)
		fmt.Println(err)
	}()
	i = 0
	err = nil
}

func TestErrEquals(t *testing.T) {
	fmt.Println(errors.New("123") == errors.New("123"))
}

// 锁是可以复用的，但是不是重入的
func TestLock(t *testing.T) {
	l := sync.Mutex{}

	// 正常
	l.Lock()
	l.Unlock()
	l.Lock()
	l.Unlock()
	// 可复用

	l.Lock()
	// 会被锁住（废话，这里在测什么东西）
	l.Lock()
	l.Unlock()
}

// 处理数位计算，在 Java 中是使用 BigDecimal，或者通用的是说，当在一定精度内相等，就认为是相等
func TestCal(t *testing.T) {
	sum1 := "0.1"
	sum2 := "0.2"
	sum3 := "0.3"

	s1, _ := strconv.ParseFloat(sum1, 64)
	s2, _ := strconv.ParseFloat(sum2, 64)
	s3, _ := strconv.ParseFloat(sum3, 64)

	fmt.Println(s1+s2 == s3)

	// 编译器提示 false，但是实际执行结果为 true
	fmt.Println(0.1+0.2 == 0.3)
}

func TestWhatever(t *testing.T) {
	a := 100 * 0.02

	// 将 100 和 0.02 拆分就不行（mismatched types int and float64）
	//a1 := 100
	//a2 := 0.02
	//a = a1 * a2

	// 猜测：像上面直接写数字，go 有帮忙做默认的类型转化 100（int） → 100（float64），但是你自己写定类型，就不行了
	println(a)
}
