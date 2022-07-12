package _1container_with_most_water

// 想了半个小时，没做出来，看了答案，确实体会到了，不要头铁，花费了一定的时间，就该去看答案
// O(n) ... 太强了，这道题一定要从两端出发的，利益最大化的动态规划思路去做
func maxArea2(height []int) int {
	var v int

	var i, j = 0, len(height) - 1
	for i < j {
		v = max(v, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return v
}
