package req

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// GetRes 请求资源文件
func GetRes(url string, query url.Values) ([]byte, error) {
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

// GetJSON 请求 json 数据
func GetJSON(url string, query url.Values, result interface{}) (err error) {
	// 前置
	if err = preHandle(url, result); err != nil {
		panic(err)
	}

	desc, ok := result.(IResponse)
	if !ok {
		desc = theDefault
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", url, query.Encode()))
	if err != nil {
		return fmt.Errorf("%s %s %w", desc.ErrTip(), "请求失败", err)
	}

	// 后置
	defer resp.Body.Close()
	return respJSONHandle(resp, desc, result)
}

// PostJSON 请求 json 数据（参数也是 json 数据）
func PostJSON(url string, query url.Values, args interface{}, result interface{}) (err error) {
	// 前置
	if err = preHandle(url, result); err != nil {
		panic(err)
	}

	desc, ok := result.(IResponse)
	if !ok {
		desc = theDefault
	}

	var bs = []byte(`{}`)
	if args != nil {
		bs, err = json.Marshal(args)
		if err != nil {
			return fmt.Errorf("%s %s %w", desc.ErrTip(), "准备请求参数失败", err)
		}
	}

	resp, err := http.Post(fmt.Sprintf("%s?%s", url, query.Encode()), ApplicationJSON, bytes.NewReader(bs))
	if err != nil {
		return fmt.Errorf("%s %s %w", desc.ErrTip(), "请求失败", err)
	}

	// 后置
	defer resp.Body.Close()
	return respJSONHandle(resp, desc, result)
}

func preHandle(url string, result interface{}) error {
	if len(url) == 0 {
		return errors.New("url must not be empty")
	}
	rv := reflect.ValueOf(result)
	if rv.Kind() != reflect.Ptr {
		return errors.New("non-pointer " + rv.Type().String())
	}
	if rv.IsNil() {
		return errors.New("nil " + rv.Type().String())
	}
	return nil
}

func respJSONHandle(resp *http.Response, desc IResponse, result interface{}) error {
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s %s %w", desc.ErrTip(), "读取响应失败", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %s %d %s", desc.ErrTip(), "http status code", resp.StatusCode, string(bs))
	}

	if contentType := resp.Header.Get(ContentType); !strings.HasPrefix(contentType, ApplicationJSON) {
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
