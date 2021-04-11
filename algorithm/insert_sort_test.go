package algorithm

import (
	"go-learn/util"
	"testing"
)

func TestInsertSort(t *testing.T) {
	intSlice := []int{1, 2, 45, 12, 11, 21, 0}

	InsertSort(intSlice)

	util.PrintSliceInfo(intSlice)
}

// 既定数组中第一个元素构成的序列是有序序列，逐步向数组后边的元素遍历，为后边的元素找到前面有序序列中适合自己的位置来保证序列还是有序序列
// 注意这里不是 找到位置 → 插入，不是分这么两步，而是一同进行的：比较交换
func InsertSort(intSlice []int) {
	times := 0

	println("插入排序开始：")
	for i := 0; i < len(intSlice); i++ {
		for j := i; j > 0 && intSlice[j] < intSlice[j-1]; j-- {
			times++
			intSlice[j], intSlice[j-1] = intSlice[j-1], intSlice[j]
		}
	}
	println("插入排序结束，判断次数：", times)
}
