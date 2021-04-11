package algorithm

import (
	"fmt"
	"testing"
)

func TestHalfSearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 7, 9, 22}

	println(HalfSearch(11, array))
}

func HalfSearch(target int, array []int) (index int, ok bool) {
	low, high := 0, len(array)
	mid := (low + high) / 2

	// 循环次数
	circleSum := 0
	for ; low <= high; mid = (low + high) / 2 {
		circleSum++

		switch {
		case target == array[mid]:
			fmt.Printf("找到了，循环了【%d】次\n", circleSum)
			return array[mid], true
		case target < array[mid]:
			high = mid - 1
		case target > array[mid]:
			low = mid + 1
		}
	}

	fmt.Printf("没找到，循环了【%d】次\n", circleSum)
	return -1, false
}
