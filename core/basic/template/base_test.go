package template

import (
	"fmt"
	"go-learn/core"
	"go-learn/util"
	"testing"
)

const tVar = `Name: {{.Name}}, Age: {{.Age}}`

func TestVar(t *testing.T) {
	data := core.Person{Age: 23, Name: "xiaoming"}

	fmt.Println(util.MustRenderString(tVar, data))
}
