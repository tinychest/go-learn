package _chain

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	chain := new(FilterChain)
	chain.Add(new(XxxFilter))
	inter := chain.Filter("1")

	fmt.Println(inter)
}
