package rpc

import (
	"fmt"
	"go-learn/core/net/rpc/dto"
	"go-learn/util/req"
)

func HTTPClient() {
	res := new(dto.Reply)

	err := req.GetJson("http://127.0.0.1:8080", nil, res)
	if err != nil {
		panic(err)
	}

	fmt.Println(*res)
}
