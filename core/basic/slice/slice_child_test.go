package slice

import (
	"go-learn/tool"
	"testing"
)

// 从一个切片划分出一个子切片，从而复用切片的内存空间地址（实际业务开发种，这里也是最容易出现内存泄漏问题）
// 语法1：slice[start:end:cap]
// - start：起始的元素下标，可以省略（默认 0）
// - end：结束的元素下标，不包括
// - cap：从起始元素开始，划分的容量，不能省略（默认为父切片的容量）
//
// 语法2：slice[start:end]
// 略去的是 cap，默认为父切片的 cap
// - start：默认 0，可以省略
// - end：默认为原切片的 len，可以省略

func TestQuota(t *testing.T) {
	affectTest(t)
	// reuseCase1(t)
	// reuseCase2(t)
}

/* 样例 */
// 当只有当 append 元素后，元素的数量大于 capacity，append 才会重新申请一个内存空间（脱离原来的切片）
func affectTest(t *testing.T) {
	ps := []int{1, 2, 3}
	cs := ps[:1:2]

	// 影响了父切片
	cs = append(cs, 4)
	tool.PrintSlice(cs)
	tool.PrintSlice(ps)

	// 没有影响父切片（超过了容量，生成了新的底层数组）
	// cs = append(cs, 4, 5)
	// util.PrintSlice(cs)
	// util.PrintSlice(ps)
}

/* 由子切片引用引申出实际开发中的切片内存空间复用 */
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
