package encode

import (
	"fmt"
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
