package _5search_insert_position

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Constraints:
//
//    1 <= nums.length <= 104
//    -104 <= nums[i] <= 104
//    nums contains distinct values sorted in ascending order.
//    -104 <= target <= 104

// if | else if | else 的提交效果比 if if if 好一些
// 因为如果目标值比 mid 对应的值小，移动的是 start，所以 start 总是能保证比目标数大
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	mid := (start + end) / 2

	for start <= end {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = (start + end) / 2
	}
	return start
}
