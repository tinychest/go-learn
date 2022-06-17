package _map

import (
	"testing"
)

type entity struct {
	Desc    string
	DescPtr *string
}

// 当通过 key 获取到 value 时，这个 value 是不可寻址的，因为 map 会进行动态扩容，当进行扩展后，map 的 value 就会进行内存迁移，其地址发生变化，所以无法对这个 value 进行寻址
//
// 因为 map 返回的值是 “不可寻址的”，所以通过 map 取得的值就只是值，也就是当前指定键对应的值，可以当作一个常量来理解，无法进行 & 的操作
//
// 《可寻址/不可寻址》在 map 中表现的简单概述：
// - 这里边会存在一个不可寻址，就是值的地址找不到，那么值就取不到的误区；说清楚这个，见这个例子
// 假如有 theMap := map[int][int]{0:0}，现在进行 &theMap[0] 获取到了 键为 0 对应的值的存储地址，之后以为这个地址就是键为 0 对应的值得地址，希望能够直接操作，
// 然后实际键 0 对应的值已经在 map 发生扩容后发生改变了，那么你以为的，就是错的了
//
// - 不可寻址强调寻到的值的地址是不确定的，换句话说取到的值是没有意义的，所以说不要进行取地址的操作，也不要进行赋值的操作，所以 Go 的编译器直接禁止了这样的操作
//
// 接着说，如果如果 map 的值是指针类型或是具有表示地址含义的（array、slice、chan、map、Ptr），虽说不能取值，但是你能通过这些地址去改变你想改变的，来满足逻辑需求
//
// 同 Java 里 testMap.get(0) = xxx 的编译不通过，说这么多其实就是你只能操作返回的值，而《不能对获取键对应值的语法，做赋值和 & 的操作》
func TestMapChangeValue(t *testing.T) {
	testChangeValue1(t)
	// testChangeValue2(t)
	// testChangeValue3(t)
}

func testChangeValue1(t *testing.T) {
	initSlice := make([]string, 3, 5)
	initSlice[0] = "0"
	initSlice[1] = "1"
	initSlice[2] = "2"
	testMap := map[int][]string{0: initSlice}

	// 修改一：正常修改
	testMap[0] = append(testMap[0], "3")
	t.Log(testMap)

	// 修改二：无论值是不是指针类型，希望通过中间变量去修改底层值
	temp := testMap[0]

	temp[0] = "-1" // 正常修改
	t.Log(testMap)

	temp = append(temp, "4")
	t.Log(len(temp))       // 5
	t.Log(len(testMap[0])) // 4
}

func testChangeValue2(t *testing.T) {
	desc := "This is a special Cup"
	changeDesc := "This is a nice Cup"
	testMap := map[string]entity{"Cup": {Desc: desc, DescPtr: &desc}}

	// 编译不通过
	// testMap["Cup"].Desc = changeDesc
	// 编译不通过（testMap["Cup"] 返回的是值类型，是指针类型就可以）
	// testMap["Cup"].DescPtr = &changeDesc
	// TODO 误区写清楚
	// TODO 可寻址 和 不可寻址，如果实在不好解释清楚，就拿更多的实例来诠释里边的概念

	// 解决一（采用指针类型）
	testMap2 := map[string]*entity{"Cup": {Desc: desc, DescPtr: &desc}}
	testMap2["Cup"].Desc = changeDesc
	testMap2["Cup"].DescPtr = &changeDesc

	// 解决二（修改替换该键对应的值）
	testMap["Cup"] = entity{Desc: changeDesc, DescPtr: &changeDesc}
}

func testChangeValue3(t *testing.T) {
	testMap := make(map[int]int, 0)
	// 这样的操作是可以的，实际效果就是为 map 添加了一对 0:1
	// 因为本质就是替换键值对
	testMap[0]++
	t.Log(testMap)
}
