package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 语法一：`json:"<字段名>[,omitempty][,string]"`
// 字段名：默认的字段名是结构体属性的字段名，而这里的值可以改变该值
// omitempty：当字段为零值时，不序列化该字段
// string：以 string 的数据类型序列化字段

// 语法二：`json:"-"`
// 序列化成 json 时，忽略该字段

// 注意：结构体的字段名首字母一定的是大写，不然序列化，反序列化，都不会考虑字段（指定了标签也没用）

// 枚举序列化 默认是 枚举 的真实值，而不是枚举的变量名（稍微想一想 type 自定义类型就应该明白）
type User struct {
	Age     int    `json:"age,omitempty,string"` // 当反序列化的 json 参数是 int，就会：json: cannot unmarshal number into Go value of type string
	Name    string `json:"na,omitempty"`
	Address string `json:"-"`
}

// 实现指定的序列化接口来修改 结构体实例 序列化的行为
// 源码核心点：encode.go/423 返回了要如何去序列化的方法，也就是 encode.go/467 → encode.go/477 调用当前方法
func (u *User) MarshalJSON() ([]byte, error) {
	u.Name = "行不改名，坐不改姓"
	return json.Marshal(u)
}

func TestJson(t *testing.T) {
	user := User{
		Age:     1,
		Name:    "小明",
		Address: "上海",
	}
	userPtr := &user

	// 序列化：默认字段名是变量名（不是默认小驼峰）
	// 如果希望输出的结果是，格式化了的，可以使用 MarshalIndent
	if userJsonStr, err := json.Marshal(user); err == nil {
		fmt.Printf("%s\n", userJsonStr)
	}

	// 反序列化：默认字母匹配就行（不区分大小写，指定了标签并不是说只按照标签匹配，如果标签没有匹配上，也是按照标签字母不区分大小写匹配的）
	// 详见：go\src\encoding\json\decode.go 704 - simpleLetterEqualFold（使用一种巧妙的 & 运算符运用，来忽略大小写的）
	if err := json.Unmarshal([]byte(`{"AgE":"11","Na":"小红"}`), userPtr); err == nil {
		fmt.Printf("%v\n", *userPtr)
	}
}
