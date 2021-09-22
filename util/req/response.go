package req

const (
	defaultSuccessCode = 0
	defaultErrTip      = ""
)

type IResponse interface {
	SuccessCode() int
	ErrTip() string
}

var defaultResponse sResponse

type sResponse struct{}

func (r sResponse) SuccessCode() int {
	return defaultSuccessCode
}

func (r sResponse) ErrTip() string {
	return defaultErrTip
}
