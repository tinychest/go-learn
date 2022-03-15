package unmarshal

import (
	"encoding/json"
	"testing"
)

type s struct {
	Name string `json:"name"`
}

func TestAddr(t *testing.T) {
	// alreadyTest(t)
	// createTest(t)
	multiNestedPtrTest(t)
}

// 创建好反序列化容器，进行反序列化
func alreadyTest(t *testing.T) {
	var theS = new(s)
	t.Logf("%p\n", theS)
	_ = json.Unmarshal([]byte(`{"name":"123"}`), theS)
	t.Logf("%p\n", theS)

	var list = make([]*s, 0, 1)
	t.Logf("%p\n", list)
	_ = json.Unmarshal([]byte(`[{"name":"123"}]`), &list)
	t.Logf("%p\n", list)
}

// 底层容器为空，让 json 类库中的代码实现创建
func createTest(t *testing.T) {
	var theS *s
	// 第二个参数直接放 theS 是不行的
	_ = json.Unmarshal([]byte(`{"name":"123"}`), &theS)
	t.Log(theS.Name)
}

// 上面可以，那就再套一层，依旧没问题
// 再套多少层都行，源码中有做一个循环解地址的操作：encoding/json/decode.go:447
func multiNestedPtrTest(t *testing.T) {
	var s1 *s
	var s2 = &s1
	var s3 = &s2
	_ = json.Unmarshal([]byte(`{"name":"123"}`), &s3)
	t.Log((*(*s3)).Name)
}
