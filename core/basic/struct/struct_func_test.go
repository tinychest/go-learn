package _struct

import (
	"testing"
)

// 函数对应的结构体，官方术语：函数 应该叫做 方法，结构体类型 应该叫做 方法的接收者（Receiver）

// 针对方法接收者的定义总结来说有如下几点（官方）
// 1、方法能够修改其接收者指向的值
// 2、可以避免在每次调用方法时复制该值（值传递）。若值的类型为大型结构体时，这样做会更加高效
// 3、通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用

// 语法点：如果方法的接收者是指针类型，我们应该说，结构体的指针类型实现了该接口，
// 而不是该结构体类型（也可以说结构体类型的方法是结构体指针类型方法的一个子集）
// Go 为什么要设计成 Xxx 和 *Xxx 都是实现了 (x Xxx) 方法，而只有 *Xxx 实现了 (x *Xxx) 方法，而 Xxx 不算，因为有些方法的实现机理就是要求改变 Receiver 的某些属性值
// （如，sync.Mutex 的 Lock 和 Unlock 方法的 Receiver 类型都是 *，否则，值拷贝将导致锁一经方法传递进行复制，就不可能实现对加的锁进行解锁了）

type st struct{}

func (s st) nor() {}

func (s *st) ptr() {}

// Invalid receiver type 'st2' ('st2' is a pointer type)
// type st2 *st
// func (s st2) abc() {}

func TestStructFunction(t *testing.T) {
	s := st{}
	sPtr := &st{}

	// 一、都可以调
	s.nor()
	s.ptr()
	sPtr.nor()
	sPtr.ptr()

	// 二、匿名有特殊
	st{}.nor()
	// st{}.ptr()  // 注意：编译不通过，匿名结构体不能进行指针类型的方法调用
	(&st{}).nor()
	(&st{}).ptr()

	// 三、特殊调用形式
	st.nor(s)
	// st.nor(sPtr) // 注意：编译不通过，特殊的调用方法不支持结构体指针类型
	// st.ptr(s)    // 注意：goland 不提示编译错误，cannot use s (type st) as type *st in argument to (*st).ptr
	// st.ptr(sPtr) // 注意：goland 提示编译错误， cannot use s (type st) as type *st in argument to (*st).ptr
	// (*st).ptr(s) // 注意：编译不通过，参数类型不对
	(*st).ptr(sPtr)

	// 四、不同的 Receiver 类型，对实际方法调用起的影响
	// 1、能不能真实影响真实值，不取决于实际调用方法的实例是否是指针类型，而是由方法的 Receiver 定义决定
	// 2、为 xxx 类型绑定的方法，*xxx 的实例类型可以调用，为 *xxx 类型绑定的方法，xxx 的实例也可以调用（语法糖，编译器会进行处理）
	// 但是有一点，呼应开头，xxx 能够调用 *xxx 类型的方法，但是 xxx 并不算实现了该方法签名对应的接口
	s.nor()
	s.ptr()
	sPtr.nor()
	sPtr.ptr()
	(&s).nor()
	(&s).ptr()
	(*sPtr).nor()
	(*sPtr).ptr()

	// 五、*XxxStruct 可以访问为 XxxStruct 定义的方法
	// 接口类型就没有这种说法，*XxxInterface 不能直接调用 XxxInterface 的方法
}
