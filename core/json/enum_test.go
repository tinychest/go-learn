package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 默认是 枚举 的真实值，而不是枚举的变量名
func TestEnumToJson(t *testing.T) {
	// 定义
	type GenderEnum int
	const (
		MALE   GenderEnum = 1
		FEMALE GenderEnum = 2
	)
	type person struct {
		Name   string
		Gender GenderEnum
	}

	// 测试
	p := person{
		Name:   "xiaomi",
		Gender: MALE,
	}
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("枚举转Json出错：%s\n", err)
	}
	println(string(bs))
}
