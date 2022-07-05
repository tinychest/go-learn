package slice

import (
	"go-learn/tool"
	"testing"
)

// 从一个切片划分出一个子切片，从而复用切片的内存空间地址（实际业务开发种，这里也是最容易出现内存泄漏问题）
// 语法1：slice[start:len:cap]
// - start：起始的元素下标，可以省略（默认 0）
// - len：从起始元素开始，划分的长度，不能省略（默认为父切片的长度）
// - cap：从起始元素开始，划分的容量，不能省略（默认为父切片的容量）
//
// 语法2：slice[start:len]
// 略去的是 cap，默认为父切片的 cap
// - start：默认 0，可以省略
// - len：默认为原切片的 len，可以省略

func TestQuota(t *testing.T) {
	affectTest(t)
	affectDetailTest(t)
}

/* 简单样例 */
func affectTest(t *testing.T) {
	var (
		ps = []int{1, 2, 3}
		cs = ps[:1:2]
	)

	// 影响了父切片
	cs = append(cs, 4)
	tool.PrintSlice(cs)
	tool.PrintSlice(ps)

	// 没有影响父切片（超过了容量，申请了新的空间）
	// cs = append(cs, 4, 5)
	// util.PrintSlice(cs)
	// util.PrintSlice(ps)
}

/* 结合内存地址，详细描述一些 */
// 其实还是 append 和 切片的原理规则，当只有当 append 元素后，元素的数量大于 capacity，append 才会重新申请一个内存空间
// 才会脱离原来的切片，没有脱离父切片的子切片，其中一方操作对应位置的元素，都会影响另外一方
func affectDetailTest(t *testing.T) {
	slice := []int{1, 2, 3}

	ref := slice[:0:0]
	// util.PrintSlice(tempSlice)
	// util.PrintSlice([]int{})
	// // 特殊的地址：0x0 length:0 capacity:0
	// util.PrintSlice(*new([]int))

	// 添加元素后切片元素个数为 1，大于切片的 capacity，所有会重新申请内存空间，创建一个新的切片
	ref = append(ref, 4)
	ref = append(ref, 5)
	ref = append(ref, 6)
	tool.PrintSlice(slice)
	tool.PrintSlice(ref)
}

/* 由子切片引用引申出实际开发中的切片内存空间复用 */
func TestReuse(t *testing.T) {
	reuseCase1(t)
	reuseCase2(t)
}

// bytes.Buffer.Reset 方法中，重置底层字节数组是这样写的：b.buf = b.buf[:0]
func reuseCase1(t *testing.T) {
	s := make([]string, 0, 4)
	s = append(s, "1")
	tool.PrintSlice(s)

	s = s[:0]
	tool.PrintSlice(s)
}

// 在 Go 中，复用切片内存因为根据使用切片的方式，可以划分成两种
// - 根据下标使用，切片使用用无需做特殊操作，特定场景如果需要严谨，复用前可以将数组的元素都置零
// - append，如果我们通过 append 重复使用一段内存空间，就需要通过子切片表达式来复用原切片（推荐）
func reuseCase2(t *testing.T) {
	s := make([]int, 0, 4)

	// 三元 不省略写法
	s = s[0:0:cap(s)]
	// s = s[1:0:cap(s)] // 没有这种写法
	tool.PrintSlice(s)
	// 三元 省略写法
	s = s[:0:cap(s)]
	tool.PrintSlice(s)
	// 二元 不省略写法
	s = s[0:0]
	tool.PrintSlice(s)
	// 二元 省略写法
	s = s[:0]
	tool.PrintSlice(s)

	// s[:] 的 len 默认是父切片的 len，不是 0
}
