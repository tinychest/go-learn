package sort

/*
【冒泡排序】

每次大循环都将最大的数沉淀到最后

[是否稳定] true

[空间复杂度] O(1)

[时间复杂度]
- 平均 O(n²)
- 最优 O(n)
- 最劣 O(n²)

[适用场景]
数组长度较小时
*/

func BubbleSort(arr []int) {
	var l = len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
