package generic

import "testing"

// Go 中 map 的键类型必须是可比较的，这个概念早就存在，现在以接口的形式落实了
// Go 中可比较的概念，从结果导向的角度来说，该类型的实例可以进行 == 和 != 操作
// 实际上，基础类型（int、float、string）以及由基础类型组合成的自定义结构体类型都是可比较的
// 上面的说法不够全面和专业，比较口语，详见该接口的注释
//
// [拓展]
// 可比较和可排序是两个不同的概念，目前并没有将可排序的概念明确定义出来
// 实验包 golang.org/x/exp/constraints 有定义可以了解一下（搜索关键字 Ordered）
type CommonMap[K comparable, V any] map[K]V // Goland 又提示错误，实际运行没有问题

func TestCommonMap(t *testing.T) {
	var m = make(CommonMap[string, int], 1)
	m["abc"] = 1
	t.Log(m)
}
