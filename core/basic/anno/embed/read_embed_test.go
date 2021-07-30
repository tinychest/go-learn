package _embed

import (
	"bytes"
	"embed"
	"fmt"
	"io/ioutil"
	"testing"
)

//go:embed static/*.txt
var allTxt embed.FS

//go:embed static/1.txt
var txtOne []byte

func TestReadOneIns(t *testing.T) {
	content, err := allTxt.ReadFile("static/1.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(content))
}

func TestReadOne(t *testing.T) {
	content, err := ioutil.ReadAll(bytes.NewReader(txtOne))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(content))
}
