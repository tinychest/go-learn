package json

import (
	"encoding/json"
	"testing"
)

type person struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Child *person `json:"child"`
}

/* 测试 json 序列化的循环引用 */
func TestRecursive(t *testing.T) {
	test1()
	//test2()
}

// 简单的循环引用
func test1() {
	var (
		p1, p2 person
		r      []byte
		err    error
	)
	p1 = person{Name: "爸爸", Child: &p2}
	p2 = person{Name: "儿子", Child: &p1}

	// json: unsupported value: encountered a cycle via *json.person
	if r, err = json.Marshal(p1); err != nil {
		panic(err)
	}
	println(string(r))
}

func test2() {
	var (
		p1, p2 *person
		r      []byte
		err    error
	)
	p1 = &person{Name: "爸爸", Child: p2}
	p2 = &person{Name: "儿子", Child: p1}

	// json: unsupported value: encountered a cycle via *json.person
	// p1.Child = p2

	if r, err = json.Marshal(p1); err != nil {
		panic(err)
	}
	println(string(r))
}
