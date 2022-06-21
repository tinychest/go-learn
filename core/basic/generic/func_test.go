package generic

import "testing"

// - 方法不支持泛型，但是可以通过 Receiver 使用的泛型来定义泛型，见下面的样例
//
// 案例：自定义数字切片类型提供了求和的方法
// - 感觉就是说 Go 的泛型像是定义时的动态，确实解决了大家反应的同样的方法，参数类型不同，就得修改类型复制一份

type slice[T int | float64] []T

func (s slice[T]) Sum() T {
	var t T
	for _, v := range s {
		t += v // Goland 提示错误，但是执行成功
	}
	return t
}

// - 不允许进行 type switch 和 类型断言
// - 可以进行反射（到了泛型，还需要用到反射，就要考虑需求是否真的需要用泛型了）
func (s *slice[T]) append(t T) {
	// Goland 没有提示，编译失败 invalid operation: cannot use type assertion on type parameter value t (variable of type T constrained by int|float64)
	// println(t.(int))
	// Goland 没有提示，编译失败 cannot use type switch on type parameter value t (variable of type T constrained by int|float64)
	// switch t.(type) {
	// case int:
	// case float64:
	// }

	*s = append(*s, t) // Goland 提示错误，执行成功
}

func TestP1(t *testing.T) {
	var i slice[int] = []int{1, 2, 3}
	t.Log(i.Sum())

	var f slice[float64] = []float64{1.1, 2.2, 3.3}
	t.Log(f.Sum())

	f.append(4.4)
	t.Log(f)
}
