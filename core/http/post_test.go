package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// 请求模拟点
// 1.URL 没有整参数
// 2.请求参数是 json
// 3.响应数据是 json
// 4.需要设置请求头（需要设置额外请求头，那么只能构建请求并通过 Client 发起请求）

// 重点方法
// bytes.NewReader(string)
// strings.NewReader([]byte)
// Post(url string, contentType string, body io.Reader) (resp *Response, err error)
// http.PostForm(url string, data url.Values) (resp *Response, err error) - 如果是提交表单（Content-Type - application/x-www-form-urlencoded）可以使用该方法
func TestPost(t *testing.T) {
	var (
		paramBytes []byte
		req        *http.Request
		resp       *http.Response
		err        error
	)

	var param map[string]string
	if paramBytes, err = json.Marshal(param); err != nil {
		t.Fatal(err)
	}

	// 1.构建请求实例
	if req, err = http.NewRequest(http.MethodPost, URL, bytes.NewReader(paramBytes)); err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json") // 设置请求头

	// 2.发起请求
	var client = new(http.Client)
	if resp, err = client.Do(req); err != nil {
		t.Fatal(err)
	}

	// 3.处理结果
	var result = new(XxxResult)
	if err = ParseJsonResponse(resp, result); err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
}
