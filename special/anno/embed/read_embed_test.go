package _embed

import (
	"embed"
	"testing"
)

//go:embed static/1.txt
var txtOne []byte // 使用 []byte 或者 string 类型接收都行

//go:embed static/*.txt
var allTxt embed.FS

func TestReadOne(t *testing.T) {
	t.Log(string(txtOne))
}

func TestReadList(t *testing.T) {
	content, err := allTxt.ReadFile("static/1.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(content))
}
