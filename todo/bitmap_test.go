package todo

import (
	"testing"
)

// 【位图】
// 背景，你有 10 个不重复的自然数需要排序，但是内存只装得下 6 个，该如何做呢
//
// [参考]
// https://studygolang.com/articles/18575?fr=sidebar
//
// [实际背景]
// 40 亿个自然数，1 个整数在 64 位的机器上使用 8B 存储，在 32 位的机器上使用 4B 存储
// 40 * 10^8* 4B = 1.6 * 10^10 B = 1.6 * 10^10 / 1024 KB
// ≈ 1.6 * 10^7 KB = 1.6 * 10^7 / 1024 MB
// ≈ 1.6 * 10^4 MB = 1.6 * 10^4 / 1024 GB
// ≈ 16 GB
// 你文件中有 10+G 的自然数数据，而内存限制 2G，该如何对文件中的数据继续排序

// BitMap 自定义位图实现
type BitMap struct {
	bs  []byte
	max int
}

// sum 要存储的一组数据中的最大的数
func NewBitMap(max int) *BitMap {
	// 最大的数有多大，就需要多大的 bit 容器来存储
	// - 最小存储单位是 byte，1 个 byte 有 8 个 bit，需要的存储空间为 sum / 8（使用更高效的位移做法）
	// - 万一没除尽，有余数，所以额外加 1 个字节
	return &BitMap{
		bs:  make([]byte, max>>3+1),
		max: max,
	}
}

func (b *BitMap) Exists(num int) bool {
	// 定位 num 位所在的字节数组中的字节的下标：0~7 在第 0 个字节中、...、n 在 n / 8 个字节中
	byteIdx := num >> 3
	theByte := b.bs[byteIdx]
	// 在这个 8 位 的字节中的第几位？会发现 n / 8 得到的余数就是这个位数，在二进制中，余数实际就是位移丢弃的数（右移了 3 位）
	bitIdx := num & 0b_0000_0111
	// 最后，我们希望获取一个字节中指定位的值，我们需要够造出对应位置的 1 进行与操作
	var binBitIdx byte = 1 << bitIdx

	// 注意，可不能写成 == 1，& 完的结果是这个字节的值，不是指定位数的值
	return theByte&binBitIdx != 0
}

func (b *BitMap) Set(num int) {
	byteIdx := num >> 3

	bitIdx := num & 0b_0000_0111
	var binBitIdx byte = 1 << bitIdx

	// 将字节指定的位置为 1
	b.bs[byteIdx] |= binBitIdx
}

func (b *BitMap) Remove(num int) {
	byteIdx := num >> 3

	bitIdx := num & 0b_0000_0111
	var binBitIdx byte = 1 << bitIdx

	// 将字节指定的位置为 0
	b.bs[byteIdx] &= ^(binBitIdx)
}

func (b *BitMap) Result() []int {
	// [调试过程]
	// [254 3 0 0 0 0 0 0 0 0 0 0 16]
	// [1111_1110 0000_0011 ...... 0001_0000]
	// 2_3_4_5_6_7_8 (8+1)_(8+2) ...... (96 + 5)
	// fmt.Println(b.bits)

	// 这里 cap 想获得正好的大小，是可以实现的，但是会影响算法的平时性能
	var res = make([]int, 0, b.max)
	for i, bs := range b.bs {
		for j := 0; j < 8; j++ {
			if bs&(1<<j) != 0 {
				res = append(res, i*8+j)
			}
		}
	}
	return res
}

/* 进行实际测试 */
var data = []int{100, 9, 8, 7, 6, 5, 4, 3, 2, 1}

func TestBitMap(t *testing.T) {
	// 内存只有 6 * 4B = 6 * 32 bit = 192 bit 的空间
	// 实际申请 (100 / 8 = 12) + 1 = 13 B = 13 * 8 bit = 104 bit 的空间
	bm := NewBitMap(100)

	for _, v := range data {
		bm.Set(v)
	}
	t.Log(bm.Result())
}
