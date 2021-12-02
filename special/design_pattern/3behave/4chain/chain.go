package _chain

// Filter 拦截器定义
type Filter interface {
	// Filter 返回是否被拦截住
	Filter(word string) bool
}

// FilterChain 拦截器链
type FilterChain struct {
	Chain []Filter
}

func (c *FilterChain) Add(f Filter) {
	c.Chain = append(c.Chain, f)
}

func (c *FilterChain) Filter(s string) bool {
	for _, v := range c.Chain {
		if v.Filter(s) {
			return true
		}
	}
	return false
}

// 拦截器实现

type XxxFilter struct {}

func (f *XxxFilter) Filter(s string) bool {
	return true
}