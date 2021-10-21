package unmarshal

import (
	"encoding/json"
	"fmt"
	"go-learn/core"
	"go-learn/util"
	"testing"
)

/*
json.Unmarshal(data []byte, v interface{})

data
- nil → panic
- ""（空串） → panic
- "{}" → no panic

v
- 非指针类型 → panic
- nil → panic

- map 的地址没变，是因为源码中的反射代码是：通过 SetMapIndex 给 map 设置元素的
- slice 的地址没变（当然，前提是你的 cap 得装的下，不然扩容，地址肯定要变的），详见：go\src\encoding\json\decode.go
- 虽然 map slice 零值是 nil，且传参行为是传地址，但不是指针类型

- 虽然初始传参要求 v（容器）不为空，但是总会存在 json 自己判断决定类型的时候
结构体字段类型如果是 interface{}，但是实际的值容器不为空，就用该值类型
基础类型就用对应的类型，其他都用 map[string]interface{}
*/
func TestUnmarshal(t *testing.T) {
	// map
	typeMapTest()

	// slice
	// typeSliceTest()

	// interface
	// typeInterfaceTest()

	// 如果接口没有定义方法，则现象同上
	// customInterfaceTest()

	// 字段类型不确定，字段值类型确定
	valueTest()
}

func typeMapTest() {
	j := `{"name":"xiaoming", "age":11}`

	var theMap map[string]interface{}
	if err := json.Unmarshal([]byte(j), &theMap); err != nil {
		panic(err)
	} else {
		fmt.Println(theMap)
	}
}

func typeSliceTest() {
	j := `[{"name":"xiaoming", "age":10}, {"name":"xiaohong", "age":11}]`

	var theSlice = make([]*core.Person, 0, 2)
	if err := json.Unmarshal([]byte(j), &theSlice); err != nil {
		panic(err)
	} else {
		util.PrintSliceInfo(theSlice)
	}
}

func typeInterfaceTest() {
	j := `{"name":"xiaoming", "age":11}`

	var r interface{}
	if err := json.Unmarshal([]byte(j), &r); err != nil {
		panic(err)
	} else {
		fmt.Println(r)
	}
}

func customInterfaceTest() {
	// I 中不定义方法 → ok，定义了 → panic
	type I interface {
		Hello()
	}

	j := `{"name":"xiaoming"}`
	i := new(I)
	if err := json.Unmarshal([]byte(j), i); err != nil {
		panic(err)
	} else {
		fmt.Println(i)
	}
}

func valueTest() {
	type S struct {
		P interface{}
	}

	j := `{"p":{"name":"xiaoming"}}`
	s := &S{
		P: new(core.Person),
	}

	fmt.Printf("%p\n", s.P)
	if err := json.Unmarshal([]byte(j), s); err != nil {
		panic(err)
	}
	fmt.Printf("%p\n", s.P)
}
