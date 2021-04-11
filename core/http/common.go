package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 要求 target 一定得是指针类型 且 不为空
func ParseJsonResponse(resp *http.Response, target interface{}) error {
	var (
		bytes []byte
		err   error
	)

	// 读取响应
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	// 解析响应（假设返回的是 json）
	if err = json.Unmarshal(bytes, target); err != nil {
		panic(err)
	}

	return nil
}
