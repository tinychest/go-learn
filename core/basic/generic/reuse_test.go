package generic

// - 语法定义点（泛型太长？泛型如何复用？ 开始为了规避 * 类型的语法歧义，使用了 interface 类型进行包装，有提到）
type Integers[T int | uint | int16 | uint16 | int32 | uint32 | int64 | uint64] []T

// 泛型复用
type Int interface {
	int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64
}

type Ints[T Int] []T

// 混合复用
type Float interface {
	float32 | float64
}

type Number interface {
	Int | Float
}

type Numbers[T Number] []T
type Numbers2[T Int | Float] []T
