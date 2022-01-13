package _chain

import (
	"testing"
)

func TestChain(t *testing.T) {
	chain := new(FilterChain)
	chain.Add(new(XxxFilter))
	inter := chain.Filter("1")

	t.Log(inter)
}
