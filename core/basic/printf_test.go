package basic

import (
	"fmt"
	"testing"
)

type Account struct {
	username string
	password string
}

// func (p *Account) Error() string {
// 	return "error..."
// }

func (p *Account) String() string {
	// 注意，不能在这里调用 fmt.Sprintf 方法 - 死递归
	// return fmt.Sprintf("%s\n", p)
	return "string..."
}

/*
详见 fmt/doc.go

《占位符》
 %p：地址
 %T：类型
 %v：对应类型的默认格式
 %s：string、[]byte、error
 %q：为所有的值添加引号
 %t：布尔类型

 %b：二进制（binary）
 %o：八进制（octal）
 %d：十进制（decimal）
 %x、%X：十六进制（hex）

 %U = U+%04X：Unicode 格式

 %e、%E：科学计数法
 %f：浮点数
 %q：单引号围绕的字符字面量

《补充符》
 [start:%]
 [-]
 - 从右边填充
 [+]
 - 结构体 添加字段名
 - 数值 正负号
 - 字符串 ASCII 编码的字符
 [#]
 - 结构体 会添加完整包路径的类型名、会添加字段名
 - 八进制 前导 0
 - 十六进制 前导 0x
 - %#p 十六进制的内存地址去掉前导 0x
 [空格、0]
 - 打印字符串或切片时，在字节之间用指定字符隔开
 [数字]
 - 填充多少个
 [end:类型标识]
*/
func TestPrintf(t *testing.T) {
	// a := &Account{
	// 	username: "xiaoming",
	// 	password: "123",
	// }

	// vsFormatTest()
	// qFormatTest()

	// +：{username:小明 password:123}
	// fmt.Printf("%+v\n", a)
	// #：main.User{username:"小明", password:"123"}
	// fmt.Printf("%#v\n", a)

	// builtInPrintTest()
}

// builtin 包下有个 println，仅支持基础类型（时间类型都不支持） - 并不推荐使用
func builtInPrintlnTest() {
	// println([]byte("123"))                   // [3/3]0xc000045f4d
	// fmt.Println([]byte("123"))               // [49 50 51]
	// fmt.Printf("%s\n", []byte("123")) // 123
	// fmt.Printf("%v\n", []byte("123")) // [49 50 51]
}

func vsFormatTest() {
	var a = Account{"小明", "123"}

	// case 实现了 Error() string：
	//     打印 Error 方法的返回值
	// case 实现了 String() string：
	//     打印 String 方法的返回值
	// default：
	//     &{小明 123}
	fmt.Println(&a)
	fmt.Printf("%v\n", &a)
	fmt.Printf("%s\n", &a)

	// default：
	//     {小明 123}
	fmt.Println(a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%s\n", a)
}

func qFormatTest() {
	theM := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": [1]string{"three"},
	}
	fmt.Printf("%s\n", theM)
	fmt.Printf("%v\n", theM)
	fmt.Printf("%q\n", theM) // 引号
}

// 动态模板
func temTemTest() {
	// Printf 格式模板中的特殊符号的转译字符是 %，如希望打印 % 字符，就得写成 %%
	// template := fmt.Sprintf("%%-%ds address:%%p length:%%v capacity:%%v\n", len(slice.([]int)))
	// println("template: ", template)
	// fmt.Printf(template, fmt.Sprint(slice), slice, len(value), cap(value))
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
