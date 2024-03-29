package encode

import (
	"go-learn/tool"
	"net/url"
	"testing"
)

func TestUrlParse(t *testing.T) {
	theUrl := `https://www.xyz.com/search?name=xiaoming&name=xiaohong&age=11`

	// Parse
	r1, err := url.Parse(theUrl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(tool.MustMarshalJSON(r1)))

	// ParseQuery
	r2, err := url.ParseQuery(theUrl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(tool.MustMarshalJSON(r2)))

	// ParseRequestURI
	r3, err := url.ParseRequestURI(theUrl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(tool.MustMarshalJSON(r3)))
}
