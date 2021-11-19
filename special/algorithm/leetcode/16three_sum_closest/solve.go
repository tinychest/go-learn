package _6three_sum_closest

// Constraints:
//
//    3 <= nums.length <= 1000
//    -1000 <= nums[i] <= 1000
//    -104 <= target <= 104

func threeSumClosest(nums []int, target int) int {
	quickSort(nums)
	return sortedThreeSumClosest(nums, target)
}

func sortedThreeSumClosest(nums []int, target int) int {
	var (
		result *int
		p1, p2 = 0, len(nums) - 1
		l1, l2 = nums[p1] + 1, nums[p2] + 1
	)

	for ; p1 < p2-1; p1++ {
		if nums[p1] == l1 {
			continue
		}
		l1 = nums[p1]

		for ; p1 < p2-1; p2-- {
			if nums[p2] == l2 {
				continue
			}
			l2 = nums[p2]

			for p3 := p1 + 1; p3 < p2; p3++ {
				temp := nums[p1] + nums[p2] + nums[p3]
				if temp == target {
					return target
				}
				if result == nil || absInt(target - temp) <= absInt(target - *result) {
					result = &temp
				}
			}
		}
		p2 = len(nums) - 1
		l2 = nums[p2] + 1
	}

	return *result
}

func absInt(a int) int {
	if a >= 0 {
		return a
	}
	return -1 * a
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
