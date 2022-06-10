package http

import (
	"go-learn/tool/req"
	"testing"
)

func TestPostJSON(t *testing.T) {
	var result map[string]interface{}
	err := req.PostJSON("http://127.0.0.1:8888/v1/red_pack/test", nil, nil, &result)
	if err != nil {
		t.Log(err)
	}
	t.Log(result)
}
