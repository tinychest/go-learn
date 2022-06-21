package generic

import "testing"

// Go 在 1.18 迎来了泛型，因此你的 Go SDK 至少得升级到 1.18，才能更好的阅读和测试下面的用例
// 详见 https://go.dev/ref/spec#Interface_types
// 参见 https://segmentfault.com/a/1190000041634906

// <基础类型>
// type Number[T int | float64] T           // Goland 没有提示，编译不通过 cannot use a type parameter as RHS in type declaration
// type Number[T int | float64] interface{} // 和泛型已经没有关系了
// 结论：基本类型不能只有类型形参

// <切片类型>
// - T 被称为类型形参
// - []T 的 [] 意味着 Slice 是切片类型
// - T 后面的类型限定，又或者说类型约束 的意思是 Slice 可以实例化为 []int 或者 []string 类型
// 不是说 Slice 实例化后既可以存储 int 也可以存储 string（简单了解 Go 内存分配也知道，这不可能实现）
type Slice[T int | string] []T

// 泛型的类型约束并不限定类型，可以是任意基础类型，可以是自定义结构体，甚至可以是接口类型

// type Slice[T *int] []T                                 // 语法错误，“T *int” 会被当成表达式
// type slice[T *int,] []T                                // 语法正确，泛型约束中只有一个类型，可以添加逗号消除歧义
// type Slice[T interface{*int} | interface{*string}] []T // 语法正确，使用 interface 包上（推荐）

// <映射>
// - 好像可以得出泛型语法 type 类型名[泛型名1 类型1 | ... | 类型n, 泛型名2 类型1 | ... | 类型n] 实际类型（切片、映射 可以使用前面定义的泛型作为元素类型）
type Map[K int64 | string, V string] map[K]V

// <结构体>
type Struct[T int | string] struct {
	Name string
	Data T
}

// <嵌套>
type Complicate[N int | float64, NS []N] map[N]NS

type Big[T int | int16 | int32 | int64] []T
type Small[T int | int64] Big[T] // 语法正确
// type Small2[T uint | uint64] Big[T] // 语法错误，不能定义超出 Big 约束的类型

// <其他>
// 匿名结构体不能使用泛型，这为单元测试测试泛型方法带来了不便

func TestGeneric(t *testing.T) {
	// 实例化 两个类型的切片
	var si Slice[int] = []int{1, 2, 3}
	var ss Slice[string] = []string{"1", "2", "3"}

	si = append(si, 4)
	ss = append(ss, "4")
	t.Log(si, ss)
}
