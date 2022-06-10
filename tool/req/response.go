package req

const (
	defaultSuccessCode = 0
	defaultErrTip      = ""
)

// IResponse 响应接口
type IResponse interface {
	SuccessCode() int
	ErrTip() string
}

// 默认实现 + 实例
type defaultResponse struct{}

var theDefault defaultResponse

func (r defaultResponse) SuccessCode() int {
	return defaultSuccessCode
}

func (r defaultResponse) ErrTip() string {
	return defaultErrTip
}
