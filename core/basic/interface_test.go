package basic

// 接口定义
type i interface {
	hello() i
}

// 等价接口
type i2 interface {
	hello() i
}

// 标准实现
type s struct{}

func (s s) hello() i {
	return s
}

// s1 并没有实现 I 接口
type s1 struct{}

func (s s1) Hello() s1 {
	return s
}

// 虽然实际调用表现上 s2 就好像实现了 i 接口，实际上并不是
type s2 struct {
	hello func() i
}

func concept() {
	var s = new(s)
	var i1 i = s
	var i2 i2 = s

	// 相同方法定义的接口，可以互相赋值（当然，大的可以赋值给小的）
	i2 = i1
	i1 = i2
}