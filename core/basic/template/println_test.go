package template

import (
	"go-learn/util"
	"testing"
)

// const tp = `{{"say:" .}}` illegal expression
const tP = `{{println "say:" .}}`

func TestPrintln(t *testing.T) {
	data := "hello"

	t.Log(util.MustRenderString(tP, data))
}
