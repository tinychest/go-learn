package json

import (
	"encoding/json"
	"testing"
)

/*
【json】
json 只支持 string 的键类型，而 json 中，对 map 的处理只支持部分烈性
详见 go/src/encoding/json/encode.go:820

【语法】
- `json:"<字段名>[,omitempty][,string]"`
参数1：默认的字段名是结构体属性的字段名，而这里的值可以改变该值
参数2：当字段为零值时，不序列化该字段（对结构体类型无效）
参数3：以 string 的数据类型序列化字段（只认 string 类型的字段进行反序列化）

- `json:"-"`
序列化 json 时，忽略该字段

【行为】
- 序列化：默认字段名是变量名，也就是大驼峰（不是默认小驼峰）
结构体的字段名首字母一定的是大写，不然序列化，反序列化，不会考虑字段（指定了标签也没用）

- 反序列化：默认字母匹配就行
不区分大小写，指定了标签并不是说只按照标签匹配，如果标签没有匹配上，也是按照标签字母不区分大小写匹配的

【拓展】
- 实现指定的序列化接口来指定 结构体实例 序列化的行为（反序列化同理）
    注意，再次调用 json.Marshal 或者 Unmarshal 方法的死递归

- 注意实例之间的互相引用，规避 (json: unsupported value: encountered a cycle via *xxx.Xxx)

【源码标记】
- MarshalJSON encode.go/423 → encode.go/467 → encode.go/477
- simpleLetterEqualFold decode.go/704
    使用一种巧妙的 & 运算符运用，来忽略大小写的
*/
type User struct {
	Age     int    `json:",omitempty"` // 当反序列化的 json 参数是 int，就会：json: cannot unmarshal number into Go value of type string
	Name    string `json:"na,omitempty"`
	Address string `json:"-"`
}

// User 类型没有实现 *User 的接口方法，*User 类型实现了 User 的接口方法
// （反过来，可不是这样， User 实现的方法，*User 也会实现）
// 这里有一个思维陷阱，即，不存在方法去改变一个 结构体实例 序列化 json 内部字段的行为
// 应该将内部字段定义为自定义类型
// - 字段类型是结构体类型：为该类型定义序列化行为，但是，这个内部字段值的自定义依赖于其他的内部字段，那就没有什么办法了，只能先处理好，再进行 json 序列化（下面取了个巧，不推荐）
// - 字段类型是基础类型：基础类型也可以按照上面取巧，也可以直接序列化自定义类型对应的基础类型，来避免死递归
func (u *User) MarshalJSON() ([]byte, error) {
	u.Name = "行不改名，坐不改姓"
	return json.Marshal(*u)
}

func TestJSON(t *testing.T) {
	user := User{
		Age:     1,
		Name:    "小明",
		Address: "上海",
	}
	// userPtr := &user

	// 如果希望输出的结果是，格式化了的，可以使用 MarshalIndent
	if userJSONStr, err := json.Marshal(&user); err == nil {
		t.Logf("%s\n", userJSONStr)
	}

	// if err := json.Unmarshal([]byte(`{"AgE":"11","Na":"小红"}`), userPtr); err == nil {
	// 	t.Logf("%v\n", *userPtr)
	// }
}
