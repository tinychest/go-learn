package fuzz

import (
	"fmt"
	fuzz "github.com/google/gofuzz"
	"testing"
)

// 官网：https://github.com/google/gofuzz
//
// 简介：gofuzz is a library for populating go objects with random values.
//
// 使用：go get github.com/google/gofuzz
//
// 这里重点介绍一下 api

func TestFuzz(t *testing.T) {
	// 为 int 赋予随机整数
	// fillNum()
	// 为 map 添加值随机的键值对
	// fillMap()

	// 为结构体中以指定比例数量的字段赋值
	// fillByChance()
	// 完全自定义，甚至可以实现自定义随机规则来支持 “枚举” 类型
	// fillTotal()

	// dvyukov/go-fuzz integration
	// 结合这个库可以实现自定义用于测试的数据
	//
	// func fillWithOrdered(data []byte) {
	// 	var dstFunc func(int)
	//
	// 	var i int
	// 	fuzz.NewFromGoFuzz(data).Fuzz(&i)
	// 	dstFunc(i)
	// }
}

func fillNum() {
	f := fuzz.New()
	var myInt int
	f.Fuzz(&myInt) // myInt gets a random value.
}

func fillMap() {
	type person struct {
		Name string
		Age  int
	}

	f := fuzz.New().NilChance(0).NumElements(1, 1)
	var myMap map[person]string
	f.Fuzz(&myMap) // myMap will have exactly one element.

	fmt.Println(myMap)
}

func fillByChance() {
	f := fuzz.New().NilChance(.5)
	var fancyStruct struct {
		A, B, C, D *string
	}
	f.Fuzz(&fancyStruct) // About half the pointers should be set.（不精确）

	fmt.Println(fancyStruct.A, fancyStruct.B, fancyStruct.C, fancyStruct.D)
}

func fillTotal() {
	type MyEnum string
	const (
		A MyEnum = "A"
		B MyEnum = "B"
	)
	type MyInfo struct {
		Type  MyEnum
		AInfo *string
		BInfo *string
	}

	f := fuzz.New().NilChance(0).Funcs(
		func(e *MyInfo, c fuzz.Continue) {
			switch c.Intn(2) {
			case 0:
				e.Type = A
				c.Fuzz(&e.AInfo)
			case 1:
				e.Type = B
				c.Fuzz(&e.BInfo)
			}
		},
	)

	var myObject MyInfo
	f.Fuzz(&myObject) // Type will correspond to whether A or B info is set.

	fmt.Println(myObject)
}
