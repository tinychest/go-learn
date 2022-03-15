package unmarshal

import (
	"encoding/json"
	"testing"
)

// 本篇案例，是从一个自定义处理指定格式的字符串逻辑的坑来的，就是为了说明一点，自定义处理 json 中 string 类型的数据时，都需要考虑到其首尾的双引号
// json 是如何反序列化原生的 string 类型，详见：encoding/json/decode.go:940（Go 1.17）
type MyString string

func (s *MyString) UnmarshalJSON(bs []byte) error {
	// json 类库里的处理逻辑相当复杂，但是一般简单处理，像这样就行了
	// if len(bs) >= 2 && bs[0] == '"' && bs[len(bs)-1] == '"' {
	// 	bs = bs[1 : len(bs)-1]
	// }
	*s = MyString(bs)
	return nil
}

type twoName struct {
	Name1 MyString
	Name2 string
}

func TestUnmarshalJSON(t *testing.T) {
	var s1 MyString
	var s2 string
	var err error

	bs := []byte(`"xiaoming"`)

	err = json.Unmarshal(bs, &s1)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(bs, &s2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(s1)
	t.Log(s2)
}

func TestUnmarshalJSONEntity(t *testing.T) {
	var n twoName
	var err error

	bs := []byte(`{"name1": "xiaoming", "name2": "xiaoming"}`)
	err = json.Unmarshal(bs, &n)
	if err != nil {
		t.Fatal(err)
	}
	bs, err = json.Marshal(n)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bs))
}
