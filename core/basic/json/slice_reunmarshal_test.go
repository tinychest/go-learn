package json

import (
	"encoding/json"
	"testing"
)

// Go 中文网 2022.02.21 每日一题
// 二次 Unmarshal 对切片类型字段的影响

// https://docs.studygolang.com/pkg/encoding/json/#Unmarshal
// 关键字："To unmarshal a JSON array into a slice"（Unmarshal 方法的注释上也有该说明）

// 要将一个 JSON 数组解码到切片（slice）中，Unmarshal 将切片长度重置为零，然后将每个元素 append 到切片中。
// 特殊情况，如果将一个空的 JSON 数组解码到一个切片中，Unmarshal 会用一个新的空切片替换该切片。

// 其实追踪源码，并不能看懂为什么 Unmarshal 方法中对切片的扩容，使用的还是原来的内存地址：
// encoding/json/decode.go:545

func TestSliceReUnmarshal(t *testing.T) {
	type S struct {
		Age   int    `json:"age"`
		Name  string `json:"name"`
		Child []int  `json:"child"`
	}

	s := S{}

	jsonStr1 := `{"age": 14, "name": "potter", "child":[1,2,3]}`
	_ = json.Unmarshal([]byte(jsonStr1), &s)
	aa := s.Child
	t.Log(aa)

	jsonStr2 := `{"age": 12, "name": "potter", "child":[3,4,5,7,8,9]}`
	_ = json.Unmarshal([]byte(jsonStr2), &s)
	t.Log(aa)
}
