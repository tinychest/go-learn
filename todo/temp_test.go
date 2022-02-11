package todo

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learn/util"
	"math"
	"testing"
	"time"
)

// TODO 类型的转换规则具体原理：Int64 有时可以直接用，有时不行
// TODO syscall.Syscall

// package 包名如果和 Go 的关键字命名相同的话，调用时，import 正常，实际引用 Go 会自动在包名前加上 “_”

func TestOne(t *testing.T) {
	s := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	// typ := reflect.TypeOf(s)
	// marshalerType := reflect.TypeOf((*json.Marshaler)(nil)).Elem()
	// b := reflect.PtrTo(typ).Implements(marshalerType)
	//
	// fmt.Println(b)

	m, _ := json.Marshal(s)
	fmt.Printf("%s", m)
}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {}

func live() People {
	var stu *Student
	return stu
}

func TestW(t *testing.T) {
	// 返回 - 值：nil，类型：*Student
	// (*Student)(nil) - 值：nil，类型：*student
	// nil - 值：nil，类型：nil
	if live() == (*Student)(nil) {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

func TestSwith(t *testing.T) {
	// interface{} 万能匹配类型；调换下面 case 的顺序，可以达到打印结果不同的效果
	var i interface{} = 1
	switch i.(type) {
	case interface{}:
		t.Log(2)
	case int:
		t.Log(1)
	}
}

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

func TestChan(t *testing.T) {
	c := make(chan int)
	close(c)

	// panic，因为 case 会尝试去做，也就是实际 case 后的语句会执行
	select {
	case c <- 1:
		fmt.Println("执行了")
	default:
		fmt.Println("不执行")
	}
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

// 接口类型，是值类型，没有地址传递的概念
func TestInterfaceWrapAddr(t *testing.T) {
	type p struct{}

	// 测试
	var p1 interface{} = p{}
	var p2 = p1
	util.PrintAddr(p1)
	util.PrintAddr(p2)
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
