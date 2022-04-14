package generic

import "testing"

// 泛型函数
// 推测语法：func 方法名[泛型名1 类型1 | ... | 类型n, 泛型名2 类型1 | ... | 类型n, ...] (参数) 返回值（参数和返回值都可以直接使用定义的泛型）
// - 匿名函数不支持泛型
// - 小结：泛型类型不能直接使用，要使用的话必须传入类型实参进行实例化
func Add[T int | float64](params ...T) T {
	var t T
	for _, v := range params {
		t += v
	}
	return t
}

func TestAdd(t *testing.T) {
	// 指定泛型
	t.Log(Add[int](1, 2))
	t.Log(Add[float64](1.1, 2.2))

	// 自动推断类型
	t.Log(Add(1, 2))
	t.Log(Add(1.1, 2.2))
}
