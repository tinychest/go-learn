package spew

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

type A struct {
	a string
	b string
	c int
}

func TestSpew(t *testing.T) {
	spew.Dump(A{})
}
