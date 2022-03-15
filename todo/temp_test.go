package todo

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"
)

// package 包名如果和 Go 的关键字命名相同的话，调用时，import 正常，实际引用 Go 会自动在包名前加上 “_”

func TestNil(t *testing.T) {
	type A struct{}

	p := func(i interface{}) {
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Interface, reflect.Slice, reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
			// 判断一个类型零值为 nil 的，当前是否是 nil 很简单；但是如果类型零值不是 nil，就麻烦了
			t.Log(v.IsNil())
		default:
			t.Log(false)
		}
	}

	p((*A)(nil))
	p("1")
}

func TestW(t *testing.T) {
	ch1 := make(chan int)
	go fmt.Println(<-ch1) // 没有闭包，当然 gg
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

func TestStr(t *testing.T) {
	// gg
	// s := "123"
	// s[0] = '0'
	// t.Log(s)

	var c chan int
	go func() { <-c }() // 永远阻塞，无论是否有发送端
	c <- 1              // 永远阻塞，无论是否有接收端
}

func TestSort(t *testing.T) {
	// sort.Slice()
	// sort.SliceStable()
	//
	// sort.Stable()
	// sort.Reverse()
	//
	// sort.Ints()
	// sort.StringSlice{}.Sort()
	// sort.Float64s()
	// sort.Float64Slice{}.Sort()
	// sort.Strings()
	// sort.StringSlice{}.Sort()
}

func TestSyscall(t *testing.T) {
	caller, file, line, ok := runtime.Caller(0)
	t.Log(caller)
	t.Log(file)
	t.Log(line)
	t.Log(ok)
}

func TestNilMap(t *testing.T) {
	// nil map 可以取值，不会引发 nil pointer panic
	var m map[string]int
	t.Log(m[""])

	t.Log(map[string]int(nil)[""])
}
