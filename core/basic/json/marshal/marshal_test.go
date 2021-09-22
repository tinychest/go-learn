package marshal

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
nil → "null"
json → 和原来的值相比较，键的值和值的开头和结尾都多了双引号（并且可以正常反序列化回来）

自定义序列化行为，需要注意，如果希望返回字符串类型，需要在 []byte 的内容补上 ""，否则，直接处理就报错
可以参见源码细节：json.encOpts.quoted、func (bits floatEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) 的最后边

不能通过为基础类型定义别名，再为别名定义自定义序列化行为的方式，来定义全局的基础类型序列化方式
*/
func TestMarshal(t *testing.T) {
	// nil
	nilTest()

	// string json
	strJsonTest()
}

func nilTest() {
	if s, err := json.Marshal(nil); err != nil {
		panic(err)
	} else {
		fmt.Println(string(s) == "null")
	}
}

func strJsonTest() {
	j := `{"name":"xiaoming", "age":11}`

	// marshal
	r, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))

	// unmarshal
	if err = json.Unmarshal(r, &j); err != nil {
		panic(err)
	}
	fmt.Println(j)
}