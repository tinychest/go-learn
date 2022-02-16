// Package req_third provides quick access to third http api（follow the normal common response structure）
package req_third

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// URL
// Method
// Head Param
// Query Param
// Body Param

// GetRes 请求资源文件
func GetRes(url string, query url.Values) ([]byte, error) {
	if len(url) == 0 {
		panic("url must not be empty")
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", url, query.Encode()))
	if err != nil {
		return nil, fmt.Errorf("%s：%w", "请求失败", err)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s：%w", "读取响应失败", err)
	}
	if len(bs) == 0 {
		return nil, fmt.Errorf("%s：%w", "响应为空", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s：%s %d %s", "请求失败", "http status code", resp.StatusCode, string(bs))
	}
	return bs, nil
}

// GetJSON 请求 json 数据
func GetJSON(url string, query url.Values, resPtr interface{}) error {
	if len(url) == 0 {
		panic("url must not be empty")
	}

	// 响应实体
	result, err := NewResponse(resPtr)
	if err != nil {
		return err
	}

	// 响应错误描述
	var desc = newDefaultDescribe()
	if successCoder, ok := resPtr.(ISuccessCode); ok {
		desc.successCode = successCoder.SuccessCode()
	}
	if errTipper, ok := resPtr.(IErrTip); ok {
		desc.errTip = errTipper.ErrTip()
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", url, query.Encode()))
	if err != nil {
		return fmt.Errorf("%s：%s %w", desc.ErrTip(), "请求失败", err)
	}

	return respJSONHandle(resp, desc, result)
}

// PostJSON 请求 json 数据（参数也是 json 数据）
func PostJSON(url string, query url.Values, args interface{}, resPtr interface{}) error {
	if len(url) == 0 {
		panic("url must not be empty")
	}

	// 响应实体
	result, err := NewResponse(resPtr)
	if err != nil {
		return err
	}

	// 响应错误描述
	var desc = newDefaultDescribe()
	if successCoder, ok := resPtr.(ISuccessCode); !ok {
		desc.successCode = successCoder.SuccessCode()
	}
	if errTipper, ok := resPtr.(IErrTip); !ok {
		desc.errTip = errTipper.ErrTip()
	}

	// 请求参数
	var bs []byte
	if args == nil {
		bs = []byte(`{}`) // 也可以在服务端使用 json.UnmarshalNullable()
	} else {
		bs, err = json.Marshal(args)
		if err != nil {
			return fmt.Errorf("%s：%s %w", desc.ErrTip(), "准备请求参数失败", err)
		}
	}

	// 请求
	resp, err := http.Post(fmt.Sprintf("%s?%s", url, query.Encode()), "application/json", bytes.NewReader(bs))
	if err != nil {
		return fmt.Errorf("%s：%s %w", desc.ErrTip(), "请求失败", err)
	}

	return respJSONHandle(resp, desc, result)
}

func respJSONHandle(resp *http.Response, desc IRespDescribe, result interface{}) error {
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s：%s %w", desc.ErrTip(), "读取响应失败", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s：%s %d %s", desc.ErrTip(), "http status code", resp.StatusCode, string(bs))
	}

	if contentType := resp.Header.Get("Content-Type"); !strings.HasPrefix(contentType, "application/json") {
		return fmt.Errorf("%s：%s %s %s", desc.ErrTip(), "响应数据不是 json 格式", contentType, string(bs))
	}

	if len(bs) == 0 {
		return fmt.Errorf("%s：%s %w", desc.ErrTip(), "响应为空", err)
	}

	if err = json.Unmarshal(bs, result); err != nil {
		return fmt.Errorf("%s：%s %s %w", desc.ErrTip(), "解析响应失败", string(bs), err)
	}
	return nil
}
