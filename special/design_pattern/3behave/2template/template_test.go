package _template

import "testing"

func TestTemplate(t *testing.T) {
	err := NewXxxSMS().Send("123", "11122223333")
	if err != nil {
		t.Fatal(err)
	}

	err = NewYyySMS().Send("123", "110")
	if err != nil {
		t.Fatal(err)
	}
}
