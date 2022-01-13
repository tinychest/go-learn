package http

import (
	"go-learn/util/req"
	"testing"
)

func TestPostJson(t *testing.T) {
	var result map[string]interface{}
	err := req.PostJson("http://127.0.0.1:8888/v1/red_pack/test", nil, nil, &result)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(result)
	}
}
