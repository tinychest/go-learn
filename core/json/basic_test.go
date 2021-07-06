package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
《语法》
1、`json:"<字段名>[,omitempty][,string]"`
参数1：默认的字段名是结构体属性的字段名，而这里的值可以改变该值
参数2：当字段为零值时，不序列化该字段
参数3：以 string 的数据类型序列化字段

2、`json:"-"`
序列化 json 时，忽略该字段

《行为》
1、序列化：默认字段名是变量名，也就是大驼峰（不是默认小驼峰）
结构体的字段名首字母一定的是大写，不然序列化，反序列化，不会考虑字段（指定了标签也没用）

2、反序列化：默认字母匹配就行
不区分大小写，指定了标签并不是说只按照标签匹配，如果标签没有匹配上，也是按照标签字母不区分大小写匹配的

《拓展》
1、实现指定的序列化接口来指定 结构体实例 序列化的行为（反序列化同理）
注意不能在实现方法中再调用 json.Marshal 或者 Unmarshal 方法，否则会造成 stack overflow

2、注意实例之间的互相引用，规避 json: unsupported value: encountered a cycle via *xxx.Xxx

3、源码标记（MarshalJSON encode.go/423 → encode.go/467 → encode.go/477）

4、源码标记（simpleLetterEqualFold decode.go/704）
使用一种巧妙的 & 运算符运用，来忽略大小写的
*/
type User struct {
	Age     int    `json:"age,omitempty,string"` // 当反序列化的 json 参数是 int，就会：json: cannot unmarshal number into Go value of type string
	Name    string `json:"na,omitempty"`
	Address string `json:"-"`
}

// 注：User 类型没有实现 *User 的接口方法，*User 类型实现了 User 的接口方法
func (u *User) MarshalJSON() ([]byte, error) {
	u.Name = "行不改名，坐不改姓"
	return json.Marshal(*u)
}

// TODO 暂时没有找到比较好的方式，来对自定义结构体类型进行特殊反序列化处理
// func (u *User) UnmarshalJSON(data []byte) error {
// 	u.Name = "行不改名，坐不改姓"
// 	return nil
// }

func TestJson(t *testing.T) {
	user := User{
		Age:     1,
		Name:    "小明",
		Address: "上海",
	}
	// userPtr := &user

	// 如果希望输出的结果是，格式化了的，可以使用 MarshalIndent
	if userJsonStr, err := json.Marshal(&user); err == nil {
		fmt.Printf("%s\n", userJsonStr)
	}

	// if err := json.Unmarshal([]byte(`{"AgE":"11","Na":"小红"}`), userPtr); err == nil {
	// 	fmt.Printf("%v\n", *userPtr)
	// }
}
