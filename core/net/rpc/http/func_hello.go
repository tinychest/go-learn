package main

import (
	"encoding/json"
	"fmt"
	"go-learn/core/net/rpc/dto"
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new request...")
	w.Header()["Content-Type"] = []string{"application/json"}

	// 获取参数
	var args = new(dto.Args)

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if len(bs) != 0 {
		err = json.Unmarshal(bs, &args)
		if err != nil {
			panic(err)
		}
	}

	// 生成响应
	reply := &dto.Reply{
		Code:    200,
		Message: "hello",
	}
	bs, err = json.Marshal(reply)
	if err != nil {
		panic(err)
	}

	var n int
	for {
		bs = bs[n:]
		n, err = w.Write(bs)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			return
		}
	}
}
