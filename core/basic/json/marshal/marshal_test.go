package marshal

import (
	"encoding/json"
	"testing"
)

/*
nil（pointer） → "null"
json（string） → 和原来的值相比较，键、值的开头和结尾都多了双引号（并且可以正常反序列化回来）

自定义序列化行为（为类型实现 json.Marshaler）
参见源码细节：json.encOpts.quoted、func (bits floatEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) 的最后边
- MarshalJSON 方法必须返回一个合法的 json 数据，不行如：nil、[]byte("")，起码：[]byte("{}")、[]byte(`{"name":"小明"}`)

不能通过为基础类型定义别名，再为别名定义自定义序列化行为的方式，来定义全局的基础类型序列化方式
*/
func TestMarshal(t *testing.T) {
	// nil
	nilTest(t)

	// string json
	strJsonTest(t)
}

func nilTest(t *testing.T) {
	if s, err := json.Marshal(nil); err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(s) == "null")
	}
}

func strJsonTest(t *testing.T) {
	j := `{"name":"xiaoming", "age":11}`

	// marshal
	r, err := json.Marshal(j)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(r))

	// unmarshal
	if err = json.Unmarshal(r, &j); err != nil {
		t.Fatal(err)
	}
	t.Log(j)
}