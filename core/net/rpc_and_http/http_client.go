package rpc_and_http

import (
	"fmt"
	"go-learn/core/net/rpc_and_http/dto"
	"go-learn/tool/req"
)

func HTTPClient() {
	res := new(dto.Reply)

	err := req.GetJSON("http://127.0.0.1:8080", nil, res)
	if err != nil {
		panic(err)
	}

	fmt.Println(*res)
}
