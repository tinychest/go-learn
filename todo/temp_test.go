package todo

import (
	"fmt"
	"go-learn/util"
	"regexp"
	"testing"
	"time"
)

/*
【point】
- *Xxx（nil）可以正常调用 *Xxx 的方法
- inner 类型不能接受 outer{inner}
- go mod why -m <module> 告诉我们为什么特定的模块在 go.mod 文件中
    go mod why gopkg.in/gomail.v2
- 在 Go 中，基础类型可以比较；结构体，如果都是基础类型，那么结构体实例可以比较
    可以通过在 结构体 中添加一个 _ [0]func() 来达到，无法比较的效果
- 可以通过在 结构体 中添加一个 _ struct{} 来避免结构体的纯值实例化方式
- httputil.DumpRequest 就像注释说的，用于服务端调试 http 请求的，封装了获取请求信息
- Go 不支持切片类型进行比较（参照自动生成的 func test，可以看到深比较 reflect.DeepEqual
    https://golang.org/ref/spec#Comparison_operators

TestIORead 为了读取请求体参数调研，Filter 将 json 参数读取出来，后边还读取的出来么
https://golangnote.com/topic/195.html（CloseReader → bs → CloseReader）
如果你不需要用它，可以考虑丢弃它，例如 io.Copy(ioutil.Discard, resp.Body)
   HTTP 客户端的传输不会重用连接，直到 body 被读完和关闭。
预读 ioutil.NopCloser(bytes.NewReader(bs))
*/

// 1、omitempty 对结构体类型无用
// 2、自定义 marshal 必须为值返回一个合法的数据
// 3、A 嵌套 B，B 嵌套 C，A 实例能够直接调用 C 的方法

func TestI(t *testing.T) {
	data := regexp.MustCompile(`\s+`).Split("a > 1", -1)
	util.PrintSlice(data)

	data = regexp.MustCompile(`\s+`).Split("a > 1", 1)
	util.PrintSlice(data)
}

type MyError struct{}

func (e MyError) Error() string { return "" }

func TestNil(t *testing.T) {
	var v *MyError = nil

	var e error = nil
	var i interface{} = nil
	fmt.Println(e == nil)
	fmt.Println(i == nil)

	e = v
	i = v
	fmt.Println(e == nil)
	fmt.Println(i == nil)
}

func TestZeroMap(t *testing.T) {
	// nil map 取值没有问题，但是存值，没有初始化，就会出问题
	var m = make(map[string]map[string]int, 0)
	v, ok := m["123"]
	if !ok {
		v = make(map[string]int, 0)
	}
	v["456"] = 1
	if !ok {
		m["123"] = v
	}

	fmt.Println(m["123"]["456"])
}

func TestTime(t *testing.T) {
	n := time.Now()
	fmt.Println(n.Unix())      // 秒
	fmt.Println(n.UnixMilli()) // 豪秒
	fmt.Println(n.UnixMicro()) // 微秒
	fmt.Println(n.UnixNano())  // 纳秒
}

func TestTimeUnit(t *testing.T) {
	var d time.Duration
	var sum int64 = 1

	// d = sum * time.Second // 这个操作不允许，不能直接和变量相乘
	d = 1 * time.Second                  // 这个可以是因为编译器做了处理
	d = time.Duration(sum) * time.Second // 正确的做法

	fmt.Println(d, sum)
}
