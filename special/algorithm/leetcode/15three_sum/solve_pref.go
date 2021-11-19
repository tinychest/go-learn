package _5three_sum

// 这道题应该是由 two sum 引申出来的，所以应该以这个为核心思想来降低时间复杂度
// 结果反向优化...
func threeSumPref(nums []int) [][]int {
	quickSort(nums)

	length := len(nums)
	if length < 3 {
		return [][]int{}
	}

	var res [][]int
	for i := 0; i < length; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		res = append(res, twoSum(nums[i+1:], -1*nums[i])...)
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	length := len(nums)
	m := make(map[int]int, length)

	var res [][]int
	for i := 0; i < length; i++ {
		m[nums[i]] = i
	}

	for i := 0; i < length; i++ {
		s := nums[i]
		if i != 0 && s == nums[i-1] {
			continue
		}
		t := target - s
		if v, ok := m[t]; ok && v != i {
			if s > t {
				s, t = t, s
			}
			res = append(res, []int{-1 * target, s, t})
			delete(m, s)
		}
	}
	return res
}
