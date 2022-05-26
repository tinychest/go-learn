package todo

import (
	"bytes"
	"fmt"
	"go-learn/util"
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

// https://go.dev/blog/maps#Key types
func TestMapKeyType(t *testing.T) {
	// 切片不可比较
	// s1 := []int{1, 2, 3}
	// s2 := []int{1, 2, 3}
	// t.Log(s1 == s2)

	// 切片不能进行 hash
	// var _ map[[]int]interface{}
}

func TestCodePos(t *testing.T) {
	// 如何输出让 Goland 控制台能够识别代码位置，且点击能跳转的日志（蓝字、下划线）
	// 没有搜索相关内容，那就断点看下实际的输出内容，这样应该就可以知道了
	// - Log 方法本身会打印出代码位置
	// t.Log("1")
	// - http 地址本是蓝色下划线可点击的
	// t.Log("http:localhost:8080/)

	// 结论：Goland 只要能索引到对应位置，就会标注成可点击的样式
	fmt.Println("temp_test.go:25")
}

func TestCompare(t *testing.T) {
	// t.Log([]int{1} == []int{1})

	var _ = [...]int{1} // [1]int
	var _ = [2]int{1}   // [2]int
}

/* funny closure */
func TestFunnyClosure(t *testing.T) {
	// 闭包问题遇到了很多了，并且很容易犯错
	// 即时写的时候考虑到了，日后再看还是容易被迷惑，所以慎用闭包，如果有必要使用，一定要注释好

	i := 1
	// 注意：闭包中修改了唯一变量的值
	// 假如这个方法接下来要被调用多次，就会出现预期外的情况
	f := func() string {
		if i != 0 {
			i = 0
			return "no zero"
		}
		return "zero"
	}
	for i := 0; i < 3; i++ {
		t.Log(f())
	}
}

/* 复用后 3 个元素的空间 */
func TestReuse(t *testing.T) {
	s := make([]int, 0, 4)
	s = append(s, 1, 2)
	util.PrintSlice(s)

	// - 没有这种写法（编译不通过）
	// s = s[1:0:cap(s)]
	// - 方式 1（晃眼）
	s = s[1:][: 0 : cap(s)-1]
	// - 方式 2
	s = s[1:]
	s = s[:0:cap(s)]

	util.PrintSlice(s)
}

/* recover nil compare */
func TestRecover(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Log("not nil")
		} else {
			t.Log("nil")
		}
	}()

	// panic(1) // not nil
	// panic(nil) // nil
	// panic([]string(nil)) // not nil
	panic("") // not nil
}

/* 了解 bytes.Buffer Truncate 方法的错用，因为 os 包下也有一个 Truncate 方法 */
func TestTruncate(t *testing.T) {
	b := bytes.Buffer{}

	b.WriteString("123")

	b.Truncate(2)
	t.Log(b.String())
}

/* 确认排序条件 */
func TestSortOrder(t *testing.T) {
	ints := []int{3, 2, 5, 1}

	// Slice 方法的第二个比较方法的参数，形参名是 less，意思是，方法默认的排序是从小到大，第 i 和 j 的元素，请你自行决定哪个比较小
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})

	t.Log(ints)
}

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
	t.Log([]string{"a", "b"})   // [a b]
	t.Log([]string{`"a b"`})    // ["a b"]
	t.Log([]string{`"a`, `b"`}) // ["a b"]

	t.Log(fmt.Sprintf("%+v", []string{"1", "2", "3"}))
	t.Log(fmt.Sprintf("%#v", []string{"1", "2", "3"}))
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
