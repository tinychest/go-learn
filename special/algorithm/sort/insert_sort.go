package sort

/*
【插入排序】

假定数组中第一个元素构成的序列是有序序列
向数组后边遍历，为后面每个元素在前面有序序列中找到适合自己的位置，来保证序列还是有序序列

[是否稳定] true

[空间复杂度] O(1)

[时间复杂度]
- 平均 O(n²)
- 最优 O(n)
- 最劣 O(n²)

[适用场景]
大部分是已排序好的
*/

func InsertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
