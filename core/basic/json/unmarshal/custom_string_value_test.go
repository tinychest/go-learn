package unmarshal

import (
	"encoding/json"
	"testing"
)

// 本篇案例，是从一个自定义处理指定格式的字符串逻辑的坑来的
// 自定义处理 json 中 string 类型的数据时，需要考虑到其首尾双引号字符的问题
// json 标准库是如何反序列化 string 类型的，详见：encoding/json/decode.go:940（Go 1.17）
type MyString string

func (s *MyString) UnmarshalJSON(bs []byte) error {
	// json 类库里的处理逻辑相当复杂，但是一般情况，简单处理，像这样就行了
	// if len(bs) >= 2 && bs[0] == '"' && bs[len(bs)-1] == '"' {
	// 	bs = bs[1 : len(bs)-1]
	// }
	*s = MyString(bs)
	return nil
}

type doubleName struct {
	Name1 MyString
	Name2 string
}

func TestUnmarshalString(t *testing.T) {
	var (
		s1 MyString
		s2 string
	)

	bs := []byte(`"xiaoming"`)

	if err := json.Unmarshal(bs, &s1); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(bs, &s2); err != nil {
		t.Fatal(err)
	}

	t.Log(s1)
	t.Log(s2)
}

func TestUnmarshalString2(t *testing.T) {
	var dou doubleName

	bs := []byte(`{"name1": "xiaoming", "name2": "xiaoming"}`)

	if err := json.Unmarshal(bs, &dou); err != nil {
		t.Fatal(err)
	}
	bs, err := json.Marshal(dou)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bs))
}
