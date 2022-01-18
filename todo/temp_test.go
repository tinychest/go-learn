package todo

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
	"go-learn/core"
	"go-learn/util"
	"math"
	"sync"
	"testing"
	"time"
)

// TODO 类型的转换规则具体原理：Int64 有时可以直接用，有时不行
// TODO syscall.Syscall

func TestOnce(t *testing.T) {
	o := sync.Once{}
	f := func(n int) {
		t.Log("print start", n)
		time.Sleep(3 * time.Second)
		t.Log("print end", n)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		t.Log("run start", 1)
		o.Do(func() {
			t.Log("do start", 1)
			f(1)
			t.Log("do end", 1)
			wg.Done()
		})
		t.Log("run end", 1)
	}()

	go func() {
		t.Log("run start", 2)
		o.Do(func() {
			t.Log("do start", 2)
			f(2)
			t.Log("do end", 2)
			wg.Done()
		})
		t.Log("run end", 2)
	}()

	wg.Wait()
}

func TestKKK(t *testing.T) {
	var k interface{} = 1

	switch k.(type) {
	case interface{}: // 万能匹配？
	}
}

// package 包名如果和 Go 的关键字命名相同的话，调用时，import 正常，实际引用 Go 会自动在包名前加上 “_”

type Name []string

func (n *Name) Append(value string) {
	*n = append(*n, value)
}

func TestAbc(t *testing.T) {
	// var n *Name // gg *nil panic
	// var n = new(Name) // ok
	var n = Name{}
	n.Append("123")
}

func TestCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		t.Log("done")
	}

	<-ctx.Done()
	t.Log("done")

	<-ctx.Done()
	t.Log("done")

	// 死锁
	// select {}

	// 死锁（c 没有关闭，也没有给值的地方）
	// select {
	// case <-c:
	// }
}

func TestAlert(t *testing.T) {
	var p interface{} = new(core.Person)
	v, ok := p.(core.Person)
	t.Log(v, ok)
}

func TestFloat642Byte(t *testing.T) {
	f := 2.33
	bs := Float642Byte(f)
	t.Log(string(bs))
	f = Byte2Float64(bs)
	t.Log(f)
}

func Float642Byte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func Byte2Float64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

func TestUnmarshalSingle(t *testing.T) {
	var bs []uint8
	bs = []byte("1.0")

	var i float64
	if err := json.Unmarshal(bs, &i); err != nil {
		t.Fatal(err)
	}
	t.Log(i)
}

func TestErr(t *testing.T) {
	// Go 标准库
	err := errors.New("MySQL error")
	err = fmt.Errorf("查询出错: %w", err)
	err = fmt.Errorf("模块出错: %w", err)
	t.Log("--------------")
	t.Log(err) // 上面拼串的结果
	t.Log("--------------")
	t.Log(errors.Unwrap(err)) // 解开一层

	// 三方类库
	err = errors.New("mysql error")
	err = errors2.Wrap(err, "查询出错")
	err = errors2.Wrap(err, "模块出错")
	t.Log("--------------")
	t.Log(err)
	// t.Log("--------------")
	// t.Log(errors2.Unwrap(err)) // 无用
	t.Log("--------------")
	t.Logf("stack trace:\n%+v\n", err) // 错误的堆栈信息
	t.Log("--------------")
	t.Log(errors2.Cause(err)) // 解到最底层（第一个错误）
}

// 接口类型，是值类型，没有地址传递的概念
func TestInterfaceWrapAddr(t *testing.T) {
	type p struct{}

	// 测试
	var p1 interface{} = p{}
	var p2 = p1
	util.PrintAddr(p1)
	util.PrintAddr(p2)
}
