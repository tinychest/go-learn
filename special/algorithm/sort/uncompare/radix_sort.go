package uncompare

// 【基数排序】
//
//
//
// [是否稳定] true
//
// [空间复杂度] O(n + k)
//
// [时间复杂度]
// - 平均 O(n * k)
// - 最优 O(n * k)
// - 最劣 O(n * k)

func RadixSort(arr []int) {
	// 数组中的最大值，为了得到应该执行位数比较的次数
	var max int
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	var maxBit int
	for t := max; t > 0; t /= 10 {
		maxBit++
	}

	// 中间数据结构
	// 非排序算法都会利用到桶的概念，这里的 ten 就是基数排序算法的桶
	var ten [10]int
	tmp := make([]int, len(arr))

	// 将每个数按照每一位进行排序，按照低位到高位的顺序
	radix := 1
	for i := 0; i < maxBit; i++ {
		// 重置
		for j := 0; j < 10; j++ {
			ten[j] = 0
		}
		// 统计指定位数的值
		for j := 0; j < len(arr); j++ {
			ten[arr[j]/radix%10]++
		}

		// 借助 上面的结果 + 辅助数组空间 对原数据按照指定位的数值进行排序
		// 虽然，借助空间进行排序的做法已经在 counting_test 中了解到，但是，另外一个精华点就在于这里了
		// 抓住 ten 所有元素之和等于目标数组的元素个数去理解下面这个循环
		for j := 1; j < 10; j++ {
			ten[j] = ten[j] + ten[j-1]
		}
		// 将 arr 按照指定位的数值进行排序，结果存储到 tmp 中
		// 注意：这里一定是逆序，因为高位得到的排序结果优先级大于地位的排序结果，相同的结果按照下面的排序算法，先来的会排在后边，算大的（可能比较难理解）
		for j := len(arr) - 1; j >= 0; j-- {
			bitValue := arr[j] / radix % 10
			bitValueSum := ten[bitValue]
			tmp[bitValueSum-1] = arr[j]

			ten[bitValue]--
		}
		// 将按照指定位数值排序的结果转存回 arr 中
		for j := 0; j < len(arr); j++ {
			arr[j] = tmp[j]
		}

		// 别忘了把基数放大，调整下次循环比较的位
		radix *= 10
	}
}
