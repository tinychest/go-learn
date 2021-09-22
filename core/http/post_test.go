package http

import (
	"fmt"
	"go-learn/util/req"
	"testing"
)

func TestPost(t *testing.T) {
	var result map[string]interface{}
	err := req.PostJson("http://127.0.0.1:8888/v1/red_pack/test", nil, nil, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
