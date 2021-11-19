package _5three_sum

// Constraints:
//
//    0 <= nums.length <= 3000
//    -105 <= nums[i] <= 105

func threeSum(nums []int) [][]int {
	quickSort(nums)
	return sortedThreeSum(nums)
}

// - for 左右各设一个指针，两个指针向中间靠拢，会因为无法做到，start 不动，end 移动两格，而丢失答案
//     for 中放两个递柜又不行 → 所以就一个递归就行了
// - 去重怎么办 → 做个判断就好
// - 答案要求以左为主怎么办 → 改一下递归的点就行了
// - 下面的答案只是没有修改这一点，返回的元素都是匹配的（只是碰巧，就遍历完去做）
// func sortedThreeSum(nums []int) [][]int {
// 	if len(nums) < 3 {
// 		return [][]int{}
// 	}
//
// 	result := make([][]int, 0)
// 	p1, p2 := 0, len(nums) - 1
//
// 	temp := -1 * (nums[p1] + nums[p2])
// 	if halfSearch(nums[p1+1:p2], temp) {
// 		result = append(result, []int{nums[p1], temp, nums[p2]})
// 	}
// 	temp = -1 * (nums[p1+1] + nums[p2])
// 	if nums[p1] != nums[p1+1] && halfSearch(nums[p1+2:p2], temp) {
// 		result = append(result, []int{nums[p1+1], temp, nums[p2]})
// 	}
// 	return append(result, sortedThreeSum(nums[p1:p2-1])...)
// }

// Runtime: 96 ms, faster than 34.38% of Go online submissions for 3Sum.
// Memory Usage: 7.6 MB, less than 57.56% of Go online submissions for 3Sum.

// - 做法太烂了，提交结果也不好，leetcode 答案要收费了，但是感觉就是这样，因为找的网上答案就是这样
func sortedThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	p1, p2 := 0, len(nums) - 1

	var temp, l1, l2 = 0, nums[p1] + 1, nums[p2] + 1
	for ; p1 < p2 - 1 && nums[p1] <= 0; p1++ {
		if nums[p1] == l1 {
			continue
		}
		l1 = nums[p1]

		for ; p1 < p2 - 1 && nums[p2] >= 0; p2-- {
			if nums[p2] == l2  {
				continue
			}
			l2 = nums[p2]

			temp = -1 * (nums[p1] + nums[p2])
			if halfSearch(nums[p1+1:p2], temp) {
				result = append(result, []int{nums[p1], temp, nums[p2]})
			}
		}
		p2 = len(nums) - 1
		l2 = nums[p2] + 1
	}
	return result
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	out := arr[0]
	start, end := 0, len(arr)-1
	for start < end {
		for arr[end] >= out && start < end {
			end--
		}
		if start < end {
			arr[start] = arr[end]
		}
		for arr[start] <= out && start < end {
			start++
		}
		if start < end {
			arr[end] = arr[start]
		}
	}
	arr[start] = out
	quickSort(arr[:start])
	quickSort(arr[start+1:])
}

func halfSearch(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}
	var start, mid, end = 0, (len(arr) - 1) / 2, len(arr) - 1
	for start <= end {
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			end = mid - 1
		}
		if arr[mid] < target {
			start = mid + 1
		}
		mid = (start + end) / 2
	}

	return false
}
