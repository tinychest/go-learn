package template

import (
	"go-learn/util"
	"testing"
)

const tRange = `{{range .list}}{{.}} {{end}}`

func TestRange(t *testing.T) {
	data := map[string]interface{}{
		"list": []string{"1", "2", "3"},
	}

	t.Log(util.MustRenderString(tRange, data))
}
