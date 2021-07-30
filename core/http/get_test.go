package http

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

const URL = "http://127.0.0.1:8080"

type XxxResult struct {
	Xxx string
}

// 详见 http 包的注释，有给出发起各种请求的例子
// 点：The client must close the response body when finished with it（client 单纯表示客户端的意思）
func TestGet(t *testing.T) {
	// 1、准备参数
	params := url.Values{
		"name": []string{"xiaoming"},
		"age":  []string{"11"},
	}
	var (
		url  = fmt.Sprintf("%s?%s", URL, params.Encode())
		resp *http.Response
		err  error
	)

	// 2、发起请求
	if resp, err = http.Get(url); err != nil {
		t.Fatal(err)
	}

	// 3、处理结果
	var result = new(XxxResult)
	if err = ParseJsonResponse(resp, result); err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
}
