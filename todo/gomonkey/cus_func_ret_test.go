package gomonkey

import (
	"github.com/agiledragon/gomonkey/v2"
	"reflect"
	"testing"
)

// 定义方法返回
//
// gomonkey.ApplyMethod 方法名必须大写，否则 “retrieve method by name failed”
// gomonkey.ApplyPrivateMethod 允许 private method

type person struct {
	name string
}

func (p *person) Hello() string {
	return p.name
}

func TestDefineFunc(t *testing.T) {
	p := &person{}
	ph := gomonkey.ApplyPrivateMethod(reflect.TypeOf(p), "Hello", func(_ *person) string {
		return "doudou"
	})
	defer ph.Reset()

	t.Log(p.Hello())
}
