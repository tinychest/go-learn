package core

import (
	"fmt"
	"html"
	"net/url"
	"testing"
	"text/template"
)

func TestVar(t *testing.T) {
	_, err := func() (int, error) {
		return -1, nil
	}()
	// 注意这里的 err 不是新的变量（废话，无法是新的变量）
	f, err := func() (float64, error) {
		return 0, nil
	}()
	// 但是放在局部初始化里边是新的的独立变量（废话，不然不就不存在）
	if err := func() error { return nil }(); err != nil {

	}

	_, _ = err, f
}

func TestUrlParse(t *testing.T) {
	theUrl := `https://www.xyz.com/search?name=xiaoming&name=xiaohong&age=11`

	r1, err := url.Parse(theUrl)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r1)

	r2, err := url.ParseQuery(theUrl)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r2) // values 类型 map[string][]string

	r3, err := url.ParseQuery(theUrl)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r3)
}

func TestValidator(t *testing.T) {
	origin := `<>`

	r1 := template.HTMLEscapeString(origin)
	r2 := template.HTMLEscaper(origin)
	r3 := html.EscapeString(origin)

	r4 := html.UnescapeString(origin)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
}
