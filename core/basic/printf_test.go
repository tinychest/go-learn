package basic

import (
	"fmt"
	"math"
	"testing"
)

// 详见 fmt/doc.go
//
// 【占位符】
// %p：地址
// %T：类型
// %v：对应类型的默认格式
// %s：string、[]byte、error
// %q：为所有的值添加引号
// %t：布尔类型
// ---
// %b：二进制（binary）
// %o：八进制（octal）
// %d：十进制（decimal）
// %x、%X：十六进制（hex）
// ---
// %U = U+%04X：Unicode 格式
// ---
// %e、%E：科学计数法
// %f：浮点数
// %q：单引号围绕的字符字面量
// ---
//
// 【补充符】
// % 和 紧跟着的类型修饰符，中间是可以添加一些补充符的
// [-]
// - 从右边填充
// [+]
// - 结构体 添加字段名
// - 数值 正负号
// - 字符串 ASCII 编码的字符
// [#]
// - 处理结构体 会添加完整包路径的类型名、会添加字段名
// - 八进制 前导 0
// - 十六进制 前导 0x
// - %#p 十六进制的内存地址去掉前导 0x
// [空格、0]
// - 打印字符串或切片时，在字节之间用指定字符隔开
// [数字]
// - 填充多少个

// 【builtin 包下有个 println】
// - 仅支持最基础的类型（时间类型都不支持）
// - 如非打印简单的 string 类型值，并不推荐使用

// 【fmt.Printf 的行为决策优先级】
// - 实现了 Error() string：
//     打印 Error 方法的返回值
// - 实现了 String() string：
//     打印 String 方法的返回值
// - 按照默认行为打印

type Account struct {
	username string
	password string
}

func (p *Account) String() string {
	// 注意，不能在这里调用 fmt.Sprintf 方法 - 死递归
	// return fmt.Sprintf("%s\n", p)
	return "string..."
}

func TestPrintf(t *testing.T) {
	classicCaseTest(t)

	// vsqFormatTest(t)
	// vsqStructFormatTest(t)
	// addSymbolTest(t)
}

func classicCaseTest(t *testing.T) {
	// 不足两位在左边补 0 保证最少两位
	fmt.Printf("%02d\n", 1)
	fmt.Printf("%02d\n", 10)

	// 右对齐的数组打印
	maxWidth := 1 + int(math.Log10(float64(10000)))
	for _, v := range []int{10, 100, 10000} {
		fmt.Printf("%*d\n", maxWidth, v)
	}

	// 不足 4 位在右边补 空格 保证最少 4 位
	fmt.Printf("%-4d%s\n", 1, "END")

	// 浮点数：主要用于限制小数位数，默认策略是四舍五入
	fmt.Printf("%.4f\n", 10.1)
}

func addSymbolTest(t *testing.T) {
	a := &Account{username: "xiaoming", password: "123"}
	// +：Error 实现 → Stringer 实现
	fmt.Printf("%+v\n", a)
	// #：main.User{username:"小明", password:"123"}
	fmt.Printf("%#v\n", a)
}

func vsqFormatTest(t *testing.T) {
	theM := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": [1]string{"three"},
	}
	fmt.Printf("%s\n", theM)
	fmt.Printf("%v\n", theM)
	fmt.Printf("%q\n", theM) // 引号
}

func vsqStructFormatTest(t *testing.T) {
	var a = Account{"小明", "123"}

	// fmt.Println(&a)
	// fmt.Printf("%v\n", &a)
	// fmt.Printf("%s\n", &a)
	// fmt.Printf("%q\n", &a)

	fmt.Println(a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%s\n", a)
	fmt.Printf("%q\n", a)
}

// Sscanf 设计的非常见简单，要求非常严格的参数规则，否则：input does not match format
// 具体规则，只能是 %s %s %d %f 这样简单的空格分割的模板串
// 并且模板定义的参数个数，要和实际接收反解析结果的参数个数相同
func sscanfTest(t *testing.T) {
	var name = "MrBBQ"
	var format = "name: %s"

	str := fmt.Sprintf("name: %s", name)
	t.Logf("模板填充 1 个参数【%s】后的结果：【%s】\n", name, str)

	var (
		param    string
		paramSum int
		err      error
	)
	if paramSum, err = fmt.Sscanf(str, format, &param); err != nil {
		t.Fatal(err)
	}
	t.Logf("【%s】根据模板反解析 %d 个参数：【%s】\n", str, paramSum, param)
}
