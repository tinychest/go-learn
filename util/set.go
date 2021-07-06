package util

type Set interface {
	Add(string)
	Exists(string) bool
	ToSlice() []string
}

type set map[string]struct{}

func NewSet(sizes ...int) Set {
	var size int
	if len(sizes) > 0 {
		size = sizes[0]
	}
	return make(set, size)
}

func (s set) Add(str string) {
	s[str] = struct{}{}
}

func (s set) Exists(key string) (exists bool) {
	// 还可以有判断是否有指定元素的方法
	_, exists = s[key]
	return
}

func (s set) ToSlice() []string {
	slice := make([]string, len(s))
	for key, _ := range s {
		slice = append(slice, key)
	}
	return slice
}
