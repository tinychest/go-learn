package encode

import (
	"fmt"
	"html"
	"html/template"
	"net/url"
	"testing"
)

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
