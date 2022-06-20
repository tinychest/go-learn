package _map

import (
	"testing"
)

// 当通过 key 获取到 value 时，这个 value 是不可寻址的，因为 map 会进行动态扩容，map 会对 value 进行内存迁移，其地址发生变化
// 为了避免使用误区，Go 的编译器直接禁止了这样的操作
//
// 因为 map 返回的值是 “不可寻址的”，不可寻址强调寻到的值的地址是不确定的，通过 map 取得的值就只是值，也就是当前指定键对应的值，可以当作一个常量来理解，无法进行 & 的操作
// 所以，修改 map 中值的操作是，先取出来 → 修改 → 设置回去
//
// 接着说，如果如果 map 的值是指针类型或是具有表示地址含义的（array、slice、chan、map、Ptr），通过这些地址去改变目标数据，是没有问题的
// 同 Java 里 testMap.get(0) = xxx 的编译不通过，说这么多其也实就是你只能操作返回的值

func TestMapChangeValue(t *testing.T) {
	type entity struct {
		Desc    string
		DescPtr *string
	}

	desc := "This is a special Cup"
	changeDesc := "This is a nice Cup"
	testMap := map[string]entity{"Cup": {Desc: desc, DescPtr: &desc}}

	// 解决一（采用指针类型）
	testMap2 := map[string]*entity{"Cup": {Desc: desc, DescPtr: &desc}}
	testMap2["Cup"].Desc = changeDesc
	testMap2["Cup"].DescPtr = &changeDesc

	// 解决二（修改替换该键对应的值）
	testMap["Cup"] = entity{Desc: changeDesc, DescPtr: &changeDesc}

	// 特例，这样的操作是可以的，实际效果就是为 map 添加了一对 0:1 TODO 通过指定参数编译，查看语法糖的本质
	m := make(map[int]int, 0)
	m[0]++
	t.Log(m)
}
