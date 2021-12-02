package req

import (
	"errors"
	"reflect"
)

const (
	DefaultSuccessCode = 200
	DefaultErrTip      = ""
)

// IResponse 响应接口定义
type IResponse interface {
	SuccessCode() int
	ErrTip() string
}

// 默认实现 + 实例
var theDefault = new(defaultResponse)

type defaultResponse struct{}

func (r defaultResponse) SuccessCode() int {
	return DefaultSuccessCode
}

func (r defaultResponse) ErrTip() string {
	return DefaultErrTip
}

// CommonResponse 通用响应结构
type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(data interface{}) (*CommonResponse, error) {
	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Ptr {
		return nil, errors.New("non-pointer " + rv.Type().String())
	}
	if rv.IsNil() {
		return nil, errors.New("nil " + rv.Type().String())
	}

	return &CommonResponse{Data: data}, nil
}
