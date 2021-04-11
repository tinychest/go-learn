package slice

import (
	"testing"
)

func TestForRangeEffectToSlice(t *testing.T) {

	// 局部定义是存在先后顺序的
	type address struct {
		name string
	}
	type person struct {
		name    string
		age     int
		address address
		values  []int
		// values interface{}
	}

	persons1 := []person{{"小明", 11, address{"北京"}, []int{0}}, {"小红", 10, address{"上海"}, []int{0}}}
	persons2 := &([]person{{"小明", 11, address{"北京"}, []int{0}}, {"小红", 10, address{"上海"}, []int{0}}})
	persons3 := []*person{{"小明", 11, address{"北京"}, []int{0}}, {"小红", 10, address{"上海"}, []int{0}}}
	persons4 := &([]*person{{"小明", 11, address{"北京"}, []int{0}}, {"小红", 10, address{"上海"}, []int{0}}})

	for _, p := range persons1 {
		p.name = "无名"
		// p.values = []int{1}
		p.values[0] = 1
	}
	for _, p := range *persons2 {
		p.name = "无名"
		// p.values = []int{1}
		p.values[0] = 1
	}
	for _, p := range persons3 {
		p.name = "无名"
		// p.values = []int{1}
		p.values[0] = 1
	}
	for _, p := range *persons4 {
		p.name = "无名"
		// p.values = []int{1}
		p.values[0] = 1
	}

	println("persons1: ", persons1)
	println("persons2: ", *persons2)
	println("persons3: ", *persons3[0], *persons3[1])
	println("persons4: ", *(*persons4)[0], *(*persons4)[1])

	// 结论：for range 结构会对要遍历的 slice（不能遍历指针类型，强行用一个指针类型，实际遍历还是要解引用，白搭）
	// 进行一次拷贝，和函数调用的参数传值的概念是一样的 - 只有指针类型，即指向了真实数据的地址，才能起到修改的效果
	// ps：切片 和 接口 在传递这个概念时，都相当于指针
	// 就是别去钻 “for range 的时候就是一个切片，不就是指针么”，这个东西就是遍历切片，且原理是会拷贝副本

	// 就是要重点提一下，不是一下蒙头进哪个哪个结构体里的字段是不是指针类型，会不会有不同的影响，而是要从一个大的全局出发，从大全局复制开始
	// 要用这样的方式才好理解，理解的全（本来就应该是这样理解，不知道自己是怎么思考的...虽然这些话也都是自己独立思考的出来的...这套娃套的）

	// 真理：从外至内，拨清复制的是值类型（基础类型、结构体类型），还是指针类型

	// 如果实在觉得复杂结构体的拷贝麻烦，就用普通的下标 for 循环遍历吧（怎么可能麻烦，得弄清楚）
}
