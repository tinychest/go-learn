package uncompare

import "container/list"

// 【桶排序】
//
// 桶排序是计数排序的升级版，假设数据均匀分布在一个范围中，基于某种映射规则、函数将数据划分成几份
// 每一份都用一个 “桶” 去装（链表）
// 通过上述可知，我们需要定义好映射规则，并预分配好桶
// 当将数据均匀划分到桶后，最后将桶中的数据排序合并起来，就是桶排序了
// （按照顺序，后边桶中的数据都大于前面桶中的数据）
//
// -这个算法特征性并不强，是一种分而治之的思想
// 大数据分解后，变为若干小数据，此时可以进行各自独立的排序（可以并行），最后将所有桶的结果连起来就可以了
//
// - 你会发现映射规则决定的，桶的数量足够大，其实这就是计数排序；当桶的数量只有一个，实质上就是桶中采用的排序算法了
//
// - 一个桶排序算法的好坏，取决于 [映射规则] 能否将 [数据] 均匀分布到各个桶中
//
// [是否稳定] true
//
// [空间复杂度] O(n + k)
//
// [时间复杂度]
// - 平均 O(n * k)
// - 最优 O(n * k)
// - 最劣 O(n²)

// BucketSort 映射规则采用简单的对数据按照数值范围、桶中数据的排序算法采用插入排序（在桶的数据插入过程就可以进行排序）
func BucketSort(arr []int) {
	bucketSum := 10

	buckets := make([]*list.List, 0, bucketSum)
	for i := 0; i < bucketSum; i++ {
		buckets = append(buckets, list.New())
	}
	for _, v := range arr {
		b := buckets[v/bucketSum]
		// 头插（特殊处理）
		if b.Front() == nil || v < b.Front().Value.(int) {
			b.PushFront(v)
			continue
		}
		// 尾插
		pre := b.Front()
		pos := b.Front()
		for pos != nil && v > pos.Value.(int) {
			pre = pos
			pos = pos.Next()
		}
		b.InsertAfter(v, pre)
	}

	p := 0
	for i := 0; i < bucketSum; i++ {
		b := buckets[i]
		for pos := b.Front(); pos != nil; pos = pos.Next() {
			arr[p] = pos.Value.(int)
			p++
		}
	}
}
