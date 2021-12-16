package template

import (
	"fmt"
	"go-learn/util"
	"testing"
)

// const tp = `{{"say:" .}}` illegal expression
const tP = `{{println "say:" .}}`

func TestPrintln(t *testing.T) {
	data := "hello"

	fmt.Println(util.MustRenderString(tP, data))
}
