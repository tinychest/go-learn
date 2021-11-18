package _adapter

// Mather 不合理的接口
type Mather interface {
	IsBigger(a, b float64) bool
}

// Math1 方法名差异
type Math1 struct{}

func (s Math1) Bigger(a, b float64) bool {
	return a > b
}

type Math1Adapter struct {
	M Mather
}

func (m Math1Adapter) IsBigger(a, b int) bool {
	return m.M.IsBigger(float64(a), float64(b))
}

// Math2 方法参数差异
type Math2 struct{}

func (s Math2) IsBigger(a, b int) bool {
	return a > b
}

type Math2Adapter struct{
	M Math2
}

func (m Math2Adapter) IsBigger(a, b float64) bool {
	return m.M.IsBigger(int(a), int(b))
}