package unmarshal

import (
	"encoding/json"
	"fmt"
	"testing"
)

type s struct {
	Name string `json:"name"`
}

func TestAddr(t *testing.T) {
	alreadyTest()
	createTest()
	multiNestedPtrTest()
}

// 创建好反序列化容器，进行反序列化
func alreadyTest() {
	var theS = new(s)
	fmt.Printf("%p\n", theS)
	_ = json.Unmarshal([]byte(`{"name":"123"}`), theS)
	fmt.Printf("%p\n", theS)

	var list = make([]*s, 0, 1)
	fmt.Printf("%p\n", list)
	_ = json.Unmarshal([]byte(`[{"name":"123"}]`), &list)
	fmt.Printf("%p\n", list)
}

// 底层容器为空，让 json 类库中的代码实现创建
func createTest() {
	var theS *s
	// 第二个参数直接放 theS 是不行的
	_ = json.Unmarshal([]byte(`{"name":"123"}`), &theS)
	fmt.Println(theS.Name)
}

// 上面可以，那就再套一层，依旧没问题
func multiNestedPtrTest() {
	var theS1 *s
	var theS2 = &theS1
	_ = json.Unmarshal([]byte(`{"name":"123"}`), &theS2)
	fmt.Println(theS1.Name)
}
