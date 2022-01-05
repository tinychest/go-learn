package req_third

// IRespDescribe 接口响应描述
type IRespDescribe interface {
	ISuccessCode
	IErrTip
}

// ISuccessCode 请求成功的 data.Code
type ISuccessCode interface {
	SuccessCode() int
}

// IErrTip 请求失败时的 错误提示
type IErrTip interface {
	ErrTip() string
}

// 实现
type respDescribe struct {
	successCode int
	errTip      string
}

func (r respDescribe) SuccessCode() int {
	return r.successCode
}

func (r respDescribe) ErrTip() string {
	return r.errTip
}

func newDefaultDescribe() respDescribe {
	return respDescribe{
		successCode: DefaultSuccessCode,
		errTip:      DefaultErrTip,
	}
}

// 默认实现 + 实例
// type defaultDescribe struct{}
//
// func (r defaultDescribe) SuccessCode() int {
// 	return DefaultSuccessCode
// }
//
// func (r defaultDescribe) ErrTip() string {
// 	return DefaultErrTip
// }
