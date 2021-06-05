package basic

import (
	"fmt"
	"go-learn/util"
	"testing"
)

type Account struct {
	username string
	password string
}

func (p *Account) Error() string {
	return "error..."
}
func (p *Account) String() string {
	return "string..."
}

func TestPrintf(t *testing.T) {
	var account = Account{"小明", "123"}

	// %p：打印出地址（标识符要匹配你参数的类型，才会打印出预期的效果）
	// fmt.Printf("%p\n", &account)

	// +：{username:小明 password:123}
	// fmt.Printf("%+v\n", account)

	// #：main.User{username:"小明", password:"123"}
	// fmt.Printf("%#v\n", account)

	// %s 和 %v 的行为
	vsFormatTest()

	// builtin 包下有个 println
	// println([]byte("123"))                   // [3/3]0xc000045f4d
	// fmt.Println([]byte("123"))               // [49 50 51]
	// fmt.Printf("%s\n", []byte("123")) // 123
	// fmt.Printf("%v\n", []byte("123")) // [49 50 51]

	// 下面这个用最形象的语句来描述的话就是 用来构建模板的模板
	// Printf 格式模板中的特殊符号的转译字符是 %，如希望打印 % 字符，就得写成 %%
	// template := fmt.Sprintf("%%-%ds address:%%p length:%%v capacity:%%v\n", len(slice.([]int)))
	// println("template: ", template)
	// fmt.Printf(template, fmt.Sprint(slice), slice, len(value), cap(value))

	util.Use(account)
}

func vsFormatTest() {
	var account = Account{"小明", "123"}

	// case 实现了 Error() string：打印 Error 方法的返回值
	// case 实现了 String() string：打印 String 方法的返回值
	// （Error 和 String Error 优先级高）
	// default：&{小明 123}
	fmt.Println(&account)
	fmt.Printf("%v\n", &account)
	fmt.Printf("%s\n", &account)

	// {小明 123}
	fmt.Println(account)
	fmt.Printf("%v\n", account)
	fmt.Printf("%s\n", account)
}

// Sscanf 设计的非常见简单，要求非常严格的参数规则，否则：input does not match format
// 具体规则，只能是 %s %s %d %f 这样简单的空格分割的模板串
// 并且模板定义的参数个数，要和实际接收反解析结果的参数个数相同
func sscanfTest() {
	var name = "MrBBQ"
	var format = "name: %s"

	str := fmt.Sprintf("name: %s", name)
	fmt.Printf("模板填充 1 个参数【%s】后的结果：【%s】\n", name, str)

	var (
		param    string
		paramSum int
		err      error
	)
	if paramSum, err = fmt.Sscanf(str, format, &param); err != nil {
		panic(err)
	}
	fmt.Printf("【%s】根据模板反解析 %d 个参数：【%s】\n", str, paramSum, param)
}
