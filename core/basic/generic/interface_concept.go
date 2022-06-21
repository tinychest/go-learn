package generic

import (
	"testing"
)

// Go 1.18 以前官方对 interface 的定义是：An interface type a method set called its interface
// 加了泛型后，很明显这个定义已经不适合了（现在是，方法集 + 泛型集（数据类型集）），现在更适合成为 a type set
//
// [空接口]
// 因为泛型的加入，也在 interface 中留下了足迹，interface 的作用更加弥足珍贵
// 因其使用频率较高，1.18 中为 interface{} 定义了一个别名 any，详见 builtin.go
//
// 一个命令替换项目中所有的 空接口 interface{} 为 any：
// 		gofmt -w -r 'interface{} -> any' ./...
//
// [comparable 接口]
// 详见 icomparable_test.go

// 为了保持语言兼容性，1.18 后接口分成了两种类型
//
// [Basic interface]
// 基本接口，接口定义中只有方法
// - 也就是 Go 1.18 之前理解的接口
//
// [General interface]
// 一般接口，接口定义中含有类型约束（type constraints）
// - 一般接口，除了含有类型，还定义了方法，在实现上判定上将更加苛刻：除了实现方法，实现方法的类型的底层类型还得在接口定义的类型范围内
// - 一般接口类型不能用来定义变量，只能用于泛型的类型约束中!（重要）
// - 带方法的一般接口不能作为类型并集的成员

// [一般接口样例]
type general interface {
	~string | ~[]rune
	len() int
}

type BS []byte

func (b BS) len() int {
	return len(b)
}

type SS string

func (s SS) len() int {
	return len(s)
}

func TestGeneral(t *testing.T) {
	//  一般接口类型不能用来定义变量，别妄想通过反射就能够突破限制
	// var _ general = BS{}
	// var _ general = SS("")
}

// [泛型接口]
type gi[T int | string] interface {
	hello() T
}

// [特殊限制1] 类型并集时，类型之间不能有相交的部分
// 编译错误 overlapping terms int and ~int
// type _ interface {
// 	~int | int
// }

// 但是是接口，即时有相交的部分，也没有问题
type _ interface {
	~int | interface{ int }
}

// [特殊限制2] 并集中不能有泛型类型（对于泛型接口来说）
// cannot embed a type parameter
// type gi[T int | string] interface {
// 	float32 | T
// }

// [特殊限制3] 接口不能直接或者间接并入自己

// [特殊限制4] 并集成员个数大于一的时候不能直接或间接并入 comparable 接口
