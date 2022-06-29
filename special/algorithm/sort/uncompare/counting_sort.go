package uncompare

// 计数排序（滑稽称 ”空间排序“）
//
// 假定我们直到待排序数组中最大的数不会太大，否则会导致申请的数组空间过大，导致崩溃
//
// [是否稳定] true
//
// [空间复杂度] O(k)
//
// [时间复杂度]
// - 平均 O(n + k)
// - 最优 O(n + k)
// - 最劣 O(n + k)

func CountingSort(arr []int) {
	var max int
	for _, v := range arr {
		if max < v {
			max = v
		}
	}

	slice := make([]bool, max+1)
	for _, v := range arr {
		slice[v] = true
	}

	var pos int
	for i, v := range slice {
		if v {
			arr[pos] = i
			pos++
		}
	}
}

// CountingPreSort 优化使用的内存空间，由原来多少个数就需要多少个 int（8 * 8 位），变为每个数只需要 1 位
func CountingPreSort(arr []int) {
	var max int
	for _, v := range arr {
		if max < v {
			max = v
		}
	}

	slice := make([]int64, max/64+1)
	for _, v := range arr {
		// 确定 v 将会落到哪个 int64 上
		b1 := v / 64
		b2 := v % 64
		// 给指定二进制位赋值
		slice[b1] |= 1 << b2
	}

	// debug
	// fmt.Printf("%064b\n", slice[0])

	var pos int
	for i, v := range slice {
		for b := 0; b < 64; b++ {
			// 注意：
			// - 运算符优先级
			// - 注意判断条件，要不 ”> 0“，要不 ”== 1<<b“，”== 1“ 错的离谱
			if v&(1<<b) > 0 {
				// 遍历每个数的每一位
				arr[pos] = 64*i + b
				pos++
			}
		}
	}
}
