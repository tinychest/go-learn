package http

import (
	"go-learn/util/req"
	"testing"
)

func TestGetJSON(t *testing.T) {
	var result map[string]interface{}
	err := req.GetJSON("http://127.0.0.1:8888/v1/red_pack/test", nil, result)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(result)
	}
}
