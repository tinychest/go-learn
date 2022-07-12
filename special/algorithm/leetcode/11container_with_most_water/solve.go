package _1container_with_most_water

// Constraints:
// n == height.length
// 2 <= n <= 10^5
// 0 <= height[i] <= 10^4

// O(n² + n)
// 提交结果：超时，尴尬...
func maxArea(height []int) int {
	var v int

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			v = max(v, (j-i)*min(height[i], height[j]))
		}
	}

	return v
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
