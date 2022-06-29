package sort

// 【归并排序】
//
// - 递归实现
//
// [是否稳定] true
//
// [空间复杂度] O(n)
//
// [时间复杂度]
// - 平均 O(nlogn)
// - 最优 O(nlogn)
// - 最劣 O(nlogn)

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	MergeSort(left)
	MergeSort(right)

	copy(arr, merge(left, right))
}

func merge(arr1, arr2 []int) []int {
	var res = make([]int, 0, len(arr1)+len(arr2))
	for len(arr1) != 0 && len(arr2) != 0 {
		if arr1[0] < arr2[0] {
			res = append(res, arr1[0])
			arr1 = arr1[1:]
		} else {
			res = append(res, arr2[0])
			arr2 = arr2[1:]
		}
	}
	for i := 0; i < len(arr1); i++ {
		res = append(res, arr1[i])
	}
	for i := 0; i < len(arr2); i++ {
		res = append(res, arr2[i])
	}
	return res
}
