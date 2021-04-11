package json

import (
	"encoding/json"
	"fmt"
	"go-learn/util"
	"testing"
)

// 《json.Unmarshal》
// 参数要求
// - 指针类型，且不为 nil （注意 map slice 虽然零值是 nil，传参行为也是传地址，但是并不是指针类型）
// map 的地址没变，是因为源码中的反射代码是：通过 SetMapIndex 给 map 设置元素的
// slice 的地址没变（当然，你的 cap 得装的下，不然扩容，地址肯定要变的），详见：go\src\encoding\json\decode.go（简单说一下，就是反射代码在使用代码参数的容器，beego 也是这样的）
// string 和原来的值相比较，键的值和值的开头和结尾都多了双引号

// 《json.Marshal》
// nil → "null"
func TestUnmarshalBasicType(t *testing.T) {
	mapJson := `{"name":"xiaoming", "age":"11"}`
	sliceJson := `[{"name":"xiaoming", "age":"10"}, {"name":"xiaohong", "age":"11"}]`

	var theMap = make(map[string]string)
	var theSlice = make([]*person, 0, 2)

	// map
	if err := json.Unmarshal([]byte(mapJson), &theMap); err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		fmt.Printf("%v\n", mapJson)
	}

	// slice
	if err := json.Unmarshal([]byte(sliceJson), &theSlice); err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		util.PrintSliceInfo(theSlice)
	}

	// string
	jsonStr := `{"name": "小明", "age": 11}`
	jsonBytes, _ := json.Marshal(jsonStr)
	fmt.Println(string(jsonBytes))
}

func TestMarshal(t *testing.T) {
	var strPtr *string
	s, _ := json.Marshal(strPtr)
	fmt.Println(s)
}
