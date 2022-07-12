package _2trap_rain_water

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
//
// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105
//
// Constraints:
//
//    n == height.length
//    1 <= n <= 2 * 104
//    0 <= height[i] <= 105

func trap(height []int) int {
	// 简单分析是说，相较于 11 题，这里更像是有多个水槽

	// return step1(height)
	return step2(height)
}

// 不愧定为 hard 的题，木有特别好的思路，先用死办法实现一个
// 实现完后，很明显，时间复杂度 O(n²)
func step1(height []int) int {
	// 需要直到当前格子能被围起来的最大高度，用这个高度减去当前柱子的高度就是当前方格能容纳的水量
	var res int
	for i := 1; i < len(height)-1; i++ {
		h := height[i]
		leftMax, rightMax := h, h

		for j := i; j >= 0; j-- {
			if leftMax < height[j] {
				leftMax = height[j]
			}
		}
		for j := i; j < len(height); j++ {
			if rightMax < height[j] {
				rightMax = height[j]
			}
		}

		res += min(leftMax, rightMax) - h
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 肯定有时间复杂度为 O(n) 的做法，核心还是要结合 11 题中的思想
// 11 题中就两块板子，你自己选，这里有多块板子，两块板子的上限的水位却决于矮的，假如在中间加了一块高的，其实水池就分为两边了，两边的一边的上限不变，另一边的上限提高了
func step2(height []int) int {
	var total, blackTotal int

	p1, p2, h1 := 0, 1, height[0]
	for p2 < len(height) {
		if height[p2] >= h1 {
			total += h1*(p2-p1-1) - blackTotal
			blackTotal = 0
			p1 = p2
			h1 = height[p2]
		} else {
			blackTotal += height[p2]
		}
		p2++
	}

	// 发现实现存在问题，因为假如 p1 是一个特别高的，那将不会有匹配上的情况，所以还需要逆向再来一下，这样的复杂度也仅仅是 O(2n)
	blackTotal = 0
	p1, p2, h1 = len(height), len(height)-1, height[len(height)-1]
	for p2 >= 0 {
		if height[p2] > h1 { // 就是上下两个等号，重复了，所以只有一个能加等号
			total += h1*(p1-p2-1) - blackTotal
			blackTotal = 0
			p1 = p2
			h1 = height[p2]
		} else {
			blackTotal += height[p2]
		}
		p2--
	}

	return total
}
