package unmarshal

import (
	"encoding/json"
	"testing"
)

// 假如要处理一匹格式不合理（有误，但符合 json 语法）的数据（先反序列化到结构实体中）
// - 通过正则、脚本将数据处理正确，再使用 json 标准库反序列化
// - 在程序代码中，自定义特定字段的反序列化行为，将不合理的数据消化掉

type elem struct {
	Name string `json:"name"`
}

type elems []elem

func (e *elems) UnmarshalJSON(src []byte) error {
	var s string
	err := json.Unmarshal(src, &s)
	if err != nil {
		return err
	}
	var res []elem
	err = json.Unmarshal(([]byte)(s), &res)
	if err != nil {
		return err
	}
	*e = res
	return nil
}

type entity struct {
	Elems *elems `json:"elems"`
	Name  string `json:"name"`
}

func TestSliceString(t *testing.T) {
	var e entity

	//language=json
	str := `{"elems": "[{\"name\": \"xiaoming\"}, {\"name\": \"xiaogang\"}]", "name": "xiaoming"}`

	if err := json.Unmarshal([]byte(str), &e); err != nil {
		t.Fatal(err)
	}
	bs, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bs))
}
