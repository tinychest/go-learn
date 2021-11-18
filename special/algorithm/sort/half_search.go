package sort

import "fmt"

func HalfSearch(target int, arr []int) (index int, ok bool) {
	low, high := 0, len(arr)
	mid := (low + high) / 2

	// 循环次数
	circleSum := 0
	for ; low <= high; mid = (low + high) / 2 {
		circleSum++

		switch {
		case target == arr[mid]:
			fmt.Printf("找到了，循环了【%d】次\n", circleSum)
			return arr[mid], true
		case target < arr[mid]:
			high = mid - 1
		case target > arr[mid]:
			low = mid + 1
		}
	}

	fmt.Printf("没找到，循环了【%d】次\n", circleSum)
	return -1, false
}
