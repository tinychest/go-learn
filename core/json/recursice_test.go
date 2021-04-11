package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/* 测试 json 序列化的循环引用 */
type person struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Child *person `json:"child"` // 结构体内部的结构体类型，编译不通过，不允许定义这样的结构体类型 - Invalid recursive type 'person'
}

func TestRecursive(t *testing.T) {
	test1()
	//test2()
}

// 简单的循环引用
func test1() {
	var (
		father, son person
		marshal     []byte
		err         error
	)
	father = person{Id: 1, Name: "爸爸", Child: &son}
	son = person{Id: 2, Name: "儿子", Child: &father}

	// json: unsupported value: encountered a cycle via *json.person
	if marshal, err = json.Marshal(father); err != nil {
		fmt.Printf("解析 json 报错了：%s\n", err)
		return
	}
	println(string(marshal))
}

// 对比上面，整一个傻乎乎的例子
func test2() {
	var (
		father, son *person
		marshal     []byte
		err         error
	)
	father = &person{Id: 1, Name: "爸爸", Child: son}
	son = &person{Id: 2, Name: "儿子", Child: father}

	// 打开这句注释，运行时报错：json: unsupported value: encountered a cycle via *json.person
	// （细品）
	// father.Child = son

	if marshal, err = json.Marshal(father); err != nil {
		fmt.Printf("解析 json 报错了：%s\n", err)
		return
	}
	println(string(marshal))
}
