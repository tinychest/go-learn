package req_third

import (
	"errors"
	"reflect"
)

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
