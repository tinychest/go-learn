package _midian_sorted_arrays

// 题目简述：两个有序整型数组，求合并后的中值
// 这个题目是真没什么难度，但是标记的是【hard】
// 考查：两个有序链表合并为一个有序链表

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	if l == 1 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		}
		if len(nums2) == 1 {
			return float64(nums2[0])
		}
	}

	nums := make([]int, 0, l)
	for p1, p2 := 0, 0; p1 != l1 || p2 != l2; {
		if p1 == l1 {
			nums = append(nums, nums2[p2:]...)
			break
		}
		if p2 == l2 {
			nums = append(nums, nums1[p1:]...)
			break
		}
		if nums1[p1] > nums2[p2] {
			nums = append(nums, nums2[p2])
			p2++
		} else {
			nums = append(nums, nums1[p1])
			p1++
		}
	}

	if l % 2 == 0 {
		return (float64(nums[l/2-1]) + float64(nums[l/2])) / 2
	} else {
		return float64(nums[l/2])
	}
}