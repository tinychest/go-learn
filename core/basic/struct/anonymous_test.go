package _struct

import (
	"fmt"
	"testing"
)

func TestAnonymous(t *testing.T) {
	var s1 *struct {
		Name    string `json:"name"`
		Address *struct {
			Street string `json:"street"`
			City   string `json:"city"`
		} `json:"address"`
	}
	var s2 *struct {
		Name    string `json:"name"`
		Address *struct {
			Street string `json:"street"`
			City   string `json:"city"`
		} `json:"address"`
	}
	// the underlying types are identical
	// s1 和 s2 是相同的类型，这里主要想强调 tag 必须也是相同的才算相同的类型（很合理，不然在某些方面的表现出现差异，也就不能称之为相同了）
	fmt.Println(s1 == s2)
}
