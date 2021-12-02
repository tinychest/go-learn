package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GetRes 请求资源文件
func GetRes(url string, query url.Values) ([]byte, error) {
	if len(url) == 0 {
		panic("url must not be empty")
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", url, query.Encode()))
	if err != nil {
		return nil, fmt.Errorf("%s %w", "请求失败", err)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "读取响应失败", err)
	}

	if len(bs) == 0 {
		return nil, fmt.Errorf("%s %w", "响应为空", err)
	}

	return bs, nil
}

// GetJson 请求 json 数据
func GetJson(url string, query url.Values, res interface{}) error {
	if len(url) == 0 {
		panic("url must not be empty")
	}

	// 响应实体
	result, err := NewResponse(res)
	if err != nil {
		return err
	}

	// 响应错误描述
	desc, ok := res.(IResponse)
	if !ok {
		desc = theDefault
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", url, query.Encode()))
	if err != nil {
		return fmt.Errorf("%s %s %w", desc.ErrTip(), "请求失败", err)
	}

	// 后置
	return respJsonHandle(resp, desc, result)
}

// PostJson 请求 json 数据（参数也是 json 数据）
func PostJson(url string, query url.Values, args interface{}, res interface{}) error {
	if len(url) == 0 {
		panic("url must not be empty")
	}

	// 响应实体
	result, err := NewResponse(res)
	if err != nil {
		return err
	}

	// 响应错误描述
	desc, ok := res.(IResponse)
	if !ok {
		desc = theDefault
	}

	// 请求参数
	var bs []byte
	if args == nil {
		bs = []byte(`{}`)
	} else {
		bs, err = json.Marshal(args)
		if err != nil {
			return fmt.Errorf("%s %s %w", desc.ErrTip(), "准备请求参数失败", err)
		}
	}

	// 请求
	resp, err := http.Post(fmt.Sprintf("%s?%s", url, query.Encode()), ApplicationJson, bytes.NewReader(bs))
	if err != nil {
		return fmt.Errorf("%s %s %w", desc.ErrTip(), "请求失败", err)
	}

	// 响应
	return respJsonHandle(resp, desc, result)
}

func respJsonHandle(resp *http.Response, desc IResponse, result interface{}) error {
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s %s %w", desc.ErrTip(), "读取响应失败", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %s %d %s", desc.ErrTip(), "http status code", resp.StatusCode, string(bs))
	}

	if contentType := resp.Header.Get(ContentType); !strings.HasPrefix(contentType, ApplicationJson) {
		return fmt.Errorf("%s %s %s %s", desc.ErrTip(), "响应数据不是 json 格式", contentType, string(bs))
	}

	if len(bs) == 0 {
		return fmt.Errorf("%s %s %w", desc.ErrTip(), "响应为空", err)
	}

	if err = json.Unmarshal(bs, result); err != nil {
		return fmt.Errorf("%s %s %s %w", desc.ErrTip(), "解析响应失败", string(bs), err)
	}
	return nil
}
