package http_rpc

import (
	"go-learn/special/performance/benchmark/http_rpc/dto"
	"go-learn/util/req"
)

func HTTPClient() {
	res := new(dto.Reply)

	err := req.GetJson("http://127.0.0.1:8080", nil, res)
	if err != nil {
		panic(err)
	}

	// log.Println(*res)
}
