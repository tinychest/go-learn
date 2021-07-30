package _struct

import (
	"fmt"
	"testing"
)

type I interface {
	hello()
}

// 含有匿名接口 且 自身实现了匿名接口
type S1 struct {I}

// 含有匿名接口 但 自身没有实现匿名接口
type S2 struct {I}

// 不含匿名接口 但实现了上面说的匿名接口
type S3 struct {}

func (h S1) hello() {
	fmt.Printf("hello one\n")
}

func (h S3) hello() {
	fmt.Printf("hello two\n")
}

func TestStructInterface(t *testing.T) {
	// 1、结构体自身有实现匿名接口的方法，具体是调用则会优先调用自身实现的方法（无论实例时，是否传入匿名接口类型的实例）
	S1{}.hello()
	S1{S3{}}.hello()

	// 2、不实现匿名接口的方法，但是可以调用（匿名包含接口的结构体被认为实现了该接口）
	// 编译通过，运行时异常：invalid memory address or nil pointer dereference
	// S2{}.hello()

	// 3、自己没倒腾出来，蒙蔽的时候，查了一下资料 https://stackoverflow.com/questions/24537443/meaning-of-a-struct-with-embedded-anonymous-interface
	// 首先，这个叫接口体匿名接口，本以为可以通过这样的方式，来达到要求结构体实例化时能够在编译时期让开发者收到错误，而实际 Go 还是秉承接口的非侵入思想（并不会编译报错提示）
	// 一个，通过这样的形式，你不用实现接口的所有方法（但是个人觉得，使用这样的设计会带来更多空指针异常的可能性）
	// 二个，在实例化一个包含匿名接口的结构体时，你可以将一个实现了匿名接口的结构体示例作为参数，那么当通过这个结构体示例（前者）调用匿名接口方法时，实际会调用参数结构体实例方法实现

	// 4、然后看了几篇博客，对 “2” 的解释：
	// 结构体内置匿名接口，实际上，底层，结构体里边会有一个对应的类型且变量名为匿名接口类型的字段，当然初始化时，不进行特殊处理，这个字段就是 nil
	// 通过结构体调用匿名接口声明的方法，实际上就是交给了这个底层字段去调用，没有处理，那肯定就是空指针
}
