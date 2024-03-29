package req_third

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

type Bg struct {
	IsDefault bool
	Code      string
	Name      string
	GroupId   int64
}

type BgData struct {
	List []Bg
}

func TestGetJSON(t *testing.T) {
	var res BgData

	err := GetJSON("", nil, &res)
	if err != nil {
		t.Fatal(err)
	}

	_, err = spew.Println(res)
	if err != nil {
		t.Fatal(err)
	}
}
