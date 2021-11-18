package sort

/*
【选择排序】

每次大循环确定最小的数，放到队首

[是否稳定] false

[空间复杂度] O(1)

[时间复杂度]
- 平均 O(n²)
- 最优 O(n²)
- 最劣 O(n²)

[适用场景]
数组长度较小时
*/

func SelectSort(arr []int) {
	var l = len(arr)
	minIndex := 0
	for i := 0; i < l; i++ {
		minIndex = i
		for j := i; j < l; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
}
