package template

import (
	"fmt"
	"go-learn/util"
	"testing"
)

const tWith = `
原来的点：{{.}}
{{- with .Word}}
现在的点：{{.}}
{{end}}`

func TestWith(t *testing.T) {
	data := map[string]interface{}{
		"Word": "WITH",
	}

	fmt.Println(util.MustRender(tWith, data))
}