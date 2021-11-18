package _map

import (
	"fmt"
	"go-learn/util"
	"testing"
)

type data struct {
	Name  string
	Slice []int
}

type item struct {
	Name  string
	Value int
}

func TestStructSlice1(_ *testing.T) {
	var m = map[string]data{
		"123": {
			Name:  "123",
			Slice: []int{1, 2, 3},
		},
	}

	// 1、同方法的值传递概念，这里就是值传递
	var (
		v1 = m["123"]
		v2 = m["123"]
	)
	fmt.Printf("%p\n", &v1)
	fmt.Printf("%p\n", &v2)

	// 2、但是，切片的地址是相同的
	var (
		s1 = m["123"].Slice
		s2 = m["123"].Slice
	)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)

	// 且改动有效
	s1[0] = 0
	fmt.Println(m)
}

func TestStructSlice2(_ *testing.T) {
	list := []item{{"a", 1}, {"a", 2}, {"b", 1}}

	var m = map[string]data{}

	for _, v := range list {
		temp, _ := m[v.Name]
		temp.Name = v.Name
		temp.Slice = append(temp.Slice, v.Value)
	}

	temp := m["a"]
	util.PrintSlice(temp.Slice)
	temp.Slice = append(temp.Slice, 1)
	util.PrintSlice(temp.Slice)
	fmt.Println(m)

	// 结论，对于 map[string]struct 的 value.field 是 slice 类型 的更新，一定得用：获取 map 的值（拷贝值） → 对拷贝进行更新 → 更新 map 的值
}

func TestStructPtrSlice(*testing.T) {
	list := []item{{"a", 1}, {"a", 2}, {"b", 1}}

	var m = map[string]*data{}

	for _, v := range list {
		temp, ok := m[v.Name]
		if !ok {
			temp = new(data)
			temp.Name = v.Name
			// 这样是可以的，因为需要关注的是初始没有实例，有了实例，本质就是一直操作这个实例
			m[v.Name] = temp
		}
		temp.Slice = append(temp.Slice, v.Value)
	}

	fmt.Printf("%#v %#v\n", m["a"], m["b"])
}
