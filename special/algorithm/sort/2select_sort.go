package sort

// 【选择排序】
//
// 每次大循环确定最小的数，放到队首
// 相较于冒泡，就是说内循环不进行交换，只有大循环切实交换位置
//
// [是否稳定] false
//
// [空间复杂度] O(1)
//
// [时间复杂度]
// - 平均 O(n²)
// - 最优 O(n²)
// - 最劣 O(n²)

func SelectSort(arr []int) {
	l := len(arr)
	minPos := 0
	for i := 0; i < l; i++ {
		minPos = i
		for j := i; j < l; j++ {
			if arr[minPos] > arr[j] {
				minPos = j
			}
		}
		if minPos != i {
			arr[minPos], arr[i] = arr[i], arr[minPos]
		}
	}
}
