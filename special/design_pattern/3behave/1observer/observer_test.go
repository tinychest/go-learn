package _observer

import "testing"

func TestSubAndPub(t *testing.T) {
	subs := Subject{}

	subs.Register(new(Observer1))
	subs.Register(new(Observer2))

	subs.Notify("123")
}
