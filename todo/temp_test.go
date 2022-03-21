package todo

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"testing"
	"time"
)

// package 包名如果和 Go 的关键字命名相同的话，调用时，import 正常，实际引用 Go 会自动在包名前加上 “_”

// 找寻后边不是跟着 bb 或 cc 的 aa "aa(?!bb|cc)"
// aabb aacc aadd
// aae aat aappo


/* nil call func，ok no nil pointer panic */
type Recorder interface {
	Record(format string, a ...interface{}) Recorder
}

type recorder []string

func (r *recorder) Record(format string, a ...interface{}) Recorder {
	if r == nil {
		*r = make([]string, 0)
	}
	*r = append(*r, fmt.Sprintf(format, a...))
	return r
}

func TestRecorder(t *testing.T) {
	var r recorder
	r.Record("haha")
	t.Log(r)
}

/* 演示业务中一个可能会忽视的排序，会影响原切片的问题 */
func TestSortAddr(t *testing.T) {
	// - 基于原切片创建新切片，修改新切片的元素值，会不会映射到老切片（取决于原切片的 cap）
	// a1 := []int{2, 1, 3}
	//
	// a1 := make([]int, 0, 4)
	// a1 = append(a1, 2, 1, 3)
	//
	// a2 := append(a1, 4)
	// a2[0] = 0
	// t.Log(a1[0])

	// - 基于原切片创建新切片，调整新切片元素位置，会不会映射到老切片
	// b1 := []int{2, 1, 4}

	b1 := make([]int, 0, 4)
	b1 = append(b1, 2, 1, 4)
	b2 := append(b1, 3)
	sort.Ints(b2)
	t.Log(b1)
}

/* interface、func call playground */
type Killer interface {
	Kill()
}

type Jack struct{}

func (j Jack) Kill() {
	fmt.Println("slash slash slash")
}

func TestMess(t *testing.T) {
	Jack{}.Kill()
	Jack.Kill(Jack{})
	interface{ Kill() }.Kill(Jack{})
	((Killer)(Jack{})).(interface{ Kill() }).Kill()
}

/* 业务中遇到的打印信息的模棱两可 */
func TestPrintPit(t *testing.T) {
	// 正常
	s := []string{"a", "b"}
	// 实际（猜测）
	s1 := []string{`"a b"`}
	// 实际（肯定）`"a,b"` → strings.Split → 如下结果
	s2 := []string{`"a`, `b"`}

	t.Log(s)  // [a b]
	t.Log(s1) // ["a b"]
	t.Log(s2) // ["a b"]
}

/* 通用的判断一个类型零值是 nil 的变量是否为 nil 的通用方法 */
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

/* 很能会忽视的，被 go 关键字迷糊的问题 */
func TestW(t *testing.T) {
	ch1 := make(chan int)
	go fmt.Println(<-ch1) // 没有闭包，当然 gg
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

/* nil chan 行为的再确认 */
func TestNilChan(t *testing.T) {
	// gg
	// s := "123"
	// s[0] = '0'
	// t.Log(s)

	var c chan int
	go func() { <-c }() // 永远阻塞，无论是否有发送端
	c <- 1              // 永远阻塞，无论是否有接收端
}

/* sort pack demo */
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

/* runtime.Caller demo */
func TestSyscall(t *testing.T) {
	caller, file, line, ok := runtime.Caller(0)
	t.Log(caller)
	t.Log(file)
	t.Log(line)
	t.Log(ok)
}

/* nil map behavior demo */
func TestNilMap(t *testing.T) {
	// nil map 可以取值，不会引发 nil pointer panic（nil slice 也可以直接 append）
	var m map[string]int
	t.Log(m[""])

	t.Log(map[string]int(nil)[""])
}
