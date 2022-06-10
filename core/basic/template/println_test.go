package template

import (
	"go-learn/tool"
	"testing"
)

// const tp = `{{"say:" .}}` illegal expression
const tP = `{{println "say:" .}}`

func TestPrintln(t *testing.T) {
	data := "hello"

	t.Log(tool.MustRenderString(tP, data))
}
