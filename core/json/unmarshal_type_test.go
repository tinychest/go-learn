package json

import (
	"encoding/json"
	"fmt"
	"go-learn/util"
	"testing"
)

/*
   《json.Unmarshal》

	参数要求：指针类型，且不为 nil （注意 map slice 虽然零值是 nil，传参行为也是传地址，但从语法类型上说并不是指针类型）

	- map 的地址没变，是因为源码中的反射代码是：通过 SetMapIndex 给 map 设置元素的
	- slice 的地址没变（当然，前提是你的 cap 得装的下，不然扩容，地址肯定要变的），详见：go\src\encoding\json\decode.go
 */
func TestDeserializeType(t *testing.T) {
	j1 := `{"name":"xiaoming", "age":11}`
	j2 := `[{"name":"xiaoming", "age":10}, {"name":"xiaohong", "age":11}]`

	var theMap = make(map[string]interface{})
	var theSlice = make([]*person, 0, 2)
	var r interface{}

	// map
	if err := json.Unmarshal([]byte(j1), &theMap); err != nil {
		panic(err)
	} else {
		fmt.Println(theMap)
	}

	// slice
	if err := json.Unmarshal([]byte(j2), &theSlice); err != nil {
		panic(err)
	} else {
		util.PrintSliceInfo(theSlice)
	}

	// interface（json 中的所有对象当作 map[string]interface{} 处理）
	if err := json.Unmarshal([]byte(j2), &r); err != nil {
		panic(err)
	} else {
		fmt.Println(r)
	}
}
