package _defer

import "testing"

// 输出：132
type S struct{}

func (s S) f(n int) S {
	print(n)
	return s
}

// 详解：
// defer 注释的函数，名为 defer (S{}).f(1).f，参数为 2
// (S{}).f(1).f 本质是 (S{}).f(1) 执行完返回的 S 结构体实例的 f 方法，所以立即就要执行，将函数地址和参数记录下来
func TestCase01(t *testing.T) {
	defer (S{}).f(1).f(2)
	print(3)
}

// 和上面的 case 一样，只是不同的实现方式
func f(n int) func(int) {
	print(n)
	return func(a int) { print(a) }
}

func TestCase02(t *testing.T) {
	defer f(1)(2)
	print(3)
}
