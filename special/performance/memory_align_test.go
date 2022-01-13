package performance

import (
	"testing"
    "unsafe"
)

/*
内存对齐（Address alignment guarantee）

一个结构体实例所占据的空间等于各字段占据空间之和，再加上内存对齐的空间大小

为什么存在内存对齐这个概念？
这个和 CPU 访问内存有关，CPU 访问内存时，以字长（word size）为单位访问
32 位的 CPU，字长位 4 字节，那么 CPU 访问内存的单位也是 4 字节
目的是加大 CPU 的吞吐量，比如同样读取 8 个字节的数据，一次读取 4 个字节，那么只需要读取 2 次，如果光是看这个没有什么特别的感觉，那看下边的小图示例

没有对齐：[1 1 1] [1 1 1] 0 0
对齐了的：[1 1 1] 0 [1 1 1] 0
假如 cpu 一次读 4 个，那么读取没有对齐的第二个数据就要读取两次

内存对齐对实现变量的原子性操作也是有好处的，每次内存访问是原子的，如果变量的大小不超过字长，那么内存对齐后，对该变量的访问就是原子的，这个特性在并发场景下至关重要。

简言之：合理的内存对齐可以提高内存读写的性能，并且便于实现变量操作的原子性。
*/

type S1 struct {
    num1 int
    num2 int
}

type S2 struct {
    num1 int16
    num2 int32
}

func TestMemoryAlignConcept(t *testing.T) {
    // 结构体实例，实际占用的空间大小
    println(unsafe.Sizeof(S1{})) // 8 + 8 = 16
    println(unsafe.Sizeof(S2{})) // 4 + 2 + 2（对齐） = 8

    // 结构体实例的 对齐系数 或叫 对齐倍数
    // unsafe.Alignof 的 规则：
    // - 对于任意类型的变量 x ，unsafe.Alignof(x) 至少为 1
    // - 对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的 unsafe.Alignof(x.f)，unsafe.Alignof(x) 等于其中的最大值
    // - 对于 array 数组类型的变量 x，unsafe.Alignof(x) 等于构成数组的元素类型的对齐倍数
    println(unsafe.Alignof(S1{})) // 8
    println(unsafe.Alignof(S2{})) // 4

    // 内存对齐的技巧
    // 一、结构体中不同类型字段的顺序，以 int8 int32 int16 为例
    // int8 int16 int32 - 8
    // [1] 0 [1 1] [1 1 1 1]
    // int8 int32 int6 - 12
    // [1] 0 0 0 [1 1 1 1] [1 1] 0 0

    // 二、空结构
    // 即当 struct{} 作为结构体最后一个字段时，需要内存对齐
    // 因为如果有指针指向该字段, 返回的地址将在结构体之外，如果此指针一直存活不释放对应的内存，就会有内存泄露的问题（该内存不因结构体释放而释放）
    // 即，struct{} 类型的字段，不能放在结构体的末尾，因为这样会花费额外的空间来保证安全
    // TODO
    type demo3 struct {
        c int32
        a struct{}
    }

    type demo4 struct {
        a struct{}
        c int32
    }

    println(unsafe.Sizeof(demo3{})) // 8
    println(unsafe.Sizeof(demo4{})) // 4
}
