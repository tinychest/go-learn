package core

import (
	"errors"
	"fmt"
	"go-learn/util"
	"html"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"text/template"
)

// type Person struct {
// 	Name    string
// 	Address *struct {
// 		Street string
// 		City   string
// 	}
// }
//
// // 匿名结构体
// var data *struct {
// 	Name    string `json:"name"`
// 	Address *struct {
// 		Street string `json:"street"`
// 		City   string `json:"city"`
// 	} `json:"address"`
// }
//
// var person = (*Person)(data)  // ignoring tags, the underlying types are identical

func TestSliceDeleteElem(t *testing.T) {
	intSlice := []int{1, 2, 3}
	err := func(intSlice *[]int, i int) error {
		if i < 0 {
			return errors.New("")
		}
		if i > len(*intSlice) {
			return errors.New("")
		}
		*intSlice = append((*intSlice)[:i-1], (*intSlice)[i:]...)
		return nil
	}(&intSlice, 1)

	if err != nil {
		panic(err)
	}

	util.PrintSliceInfo(intSlice)
}

func TestUrlParse(t *testing.T) {
	theUrl := `https://www.xyz.com/search?name=xiaoming&name=xiaohong&age=11`

	if values, err := url.Parse(theUrl); err != nil {
		panic(err)
	} else {
		fmt.Println(values)
	}

	if values, err := url.ParseQuery(theUrl); err != nil {
		panic(err)
	} else {
		// values 类型 map[string][]string
		fmt.Println(values)
	}

	if result, err := url.ParseQuery(theUrl); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

}

const sum = 0.1

func TestPrintType(t *testing.T) {
	fmt.Println(1e-1)

	typ := reflect.TypeOf(sum)
	fmt.Println(typ)

	fmt.Println(0.8999999999+1e-9 >= 0.9)
}

func TestValidator(t *testing.T) {
	origin := `<>`

	r1 := template.HTMLEscapeString(origin)
	r2 := template.HTMLEscaper(origin)
	r3 := html.EscapeString(origin)

	r4 := html.UnescapeString(origin)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
}

// defer 定义的位置，并不是去闭包的死值，是动态的（defer 传形参执行的函数是值，理会好闭包就是取，传参就是值传递的规程，就不需要解释什么）
// 两个 error 实例比较，并不是内容相同就是相等的，详见 go 比较规则
// 锁是可以复用的（废话）

// 点1：defer 的顺序（代码由上至下，压入堆栈）
// 点2：defer 闭包才能达到，最终值改变的效果
func TestDeffer(t *testing.T) {
	var i = 1

	// 最后执行：1（立即故值，延迟执行）
	defer fmt.Println("zero", i)
	// 后执行：2
	defer func() { fmt.Println("one", i) }()
	// 先执行：1（立即故值，延迟执行）
	defer func(param int) { fmt.Println("two", param) }(i)

	i = 2
}

// 处理数位计算，在 Java 中是使用 BigDecimal，或者通用的是说，当在一定精度内相等，就认为是相等
func TestCal(t *testing.T) {
	sum1 := "0.1"
	sum2 := "0.2"
	sum3 := "0.3"

	s1, _ := strconv.ParseFloat(sum1, 64)
	s2, _ := strconv.ParseFloat(sum2, 64)
	s3, _ := strconv.ParseFloat(sum3, 64)

	// 编译其不提示，但是实际执行结果为 false
	fmt.Println(s1+s2 == s3)

	// 编译器提示 false，但是实际执行结果为 true
	fmt.Println(0.1+0.2 == 0.3)
}

func TestWhatever(t *testing.T) {
	a := 100 * 0.02

	// 将 100 和 0.02 拆分就不行（mismatched types int and float64）
	// a1 := 100
	// a2 := 0.02
	// a = a1 * a2

	// 猜测：像上面直接写数字，go 有帮忙做默认的类型转化 100（int） → 100（float64），但是你自己写定类型，就不行了
	println(a)
}
