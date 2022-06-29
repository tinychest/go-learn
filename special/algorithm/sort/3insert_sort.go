package sort

// 【插入排序】
//
// 假定数组中第一个元素构成的序列是有序序列
// 向后边遍历目标数组，为有序序列添加一个元素，随后为新元素找到适合自己的位置，来保证序列还是有序序列
//
// [是否稳定] true
//
// [空间复杂度] O(1)
//
// [时间复杂度]
// - 平均 O(n²)
// - 最优 O(n)
// - 最劣 O(n²)

func InsertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// InsertSort2 交换=两次赋值操作，下面这种做法就是为目标找到合适的位置，每次只是将元素向后移，性能比上面稍微好一些
func InsertSort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i

		tmp := arr[i]
		for j > 0 && arr[j-1] > tmp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = tmp

	}
}
