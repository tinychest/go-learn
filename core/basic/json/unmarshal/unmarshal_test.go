package unmarshal

import (
	"encoding/json"
	"fmt"
	"go-learn/core"
	"go-learn/util"
	"testing"
)

/*
参数要求：指针类型，且不为 nil（注意 map slice 虽然零值是 nil，传参行为也是传地址，但从语法类型上说并不是指针类型）

- map 的地址没变，是因为源码中的反射代码是：通过 SetMapIndex 给 map 设置元素的
- slice 的地址没变（当然，前提是你的 cap 得装的下，不然扩容，地址肯定要变的），详见：go\src\encoding\json\decode.go
*/
func TestUnmarshal(t *testing.T) {
	// map
	mapTest()
	// slice
	sliceTest()
	// interface（json 中的所有对象当作 map[string]interface{} 处理）
	interfaceTest()
}

func mapTest() {
	j := `{"name":"xiaoming", "age":11}`

	var theMap = make(map[string]interface{})
	if err := json.Unmarshal([]byte(j), &theMap); err != nil {
		panic(err)
	} else {
		fmt.Println(theMap)
	}
}

func sliceTest() {
	j := `[{"name":"xiaoming", "age":10}, {"name":"xiaohong", "age":11}]`

	var theSlice = make([]*core.Person, 0, 2)
	if err := json.Unmarshal([]byte(j), &theSlice); err != nil {
		panic(err)
	} else {
		util.PrintSliceInfo(theSlice)
	}
}

func interfaceTest() {
	j := `{"name":"xiaoming", "age":11}`

	var r interface{}
	if err := json.Unmarshal([]byte(j), &r); err != nil {
		panic(err)
	} else {
		fmt.Println(r)
	}
}
