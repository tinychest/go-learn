package _52maximum_product_subarray

// Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.
//
// The test cases are generated so that the answer will fit in a 32-bit integer.
//
// A subarray is a contiguous subsequence of the array.
//
//
//
// Constraints:
//
//    1 <= nums.length <= 2 * 104
//    -10 <= nums[i] <= 10
//    The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// [理解]
// - 这里的 product 翻译成乘积
// - subarray 是指子集，不是什么连续的整数（不知道怎么会理解成这样）
//
// 所有情况全部遍历显然是不行的，因为测试用例的数组长度达到一定的数值
// 0 特殊，一旦取子集包含了 0，这个子集的结果最大就是 0 了，所以以 0 为边界划分成多个子区间
// 然后，就是根据区间中负数的个数，做出相应的处理
// 小结：断点调试了好久，修正了很多细节，要是实际面试遇到这样的题，不说做的来做不来，调试花费的时间肯定就 g 了
//
// [参考]
// 仔细想想是有在一个循环中的稍微简单一些的做法

func maxProduct(nums []int) int {
	res := -11

	// 跳过初始的 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// 不好看的写法
			res = 0
			continue
		}
		nums = nums[i:]
		break
	}
	// 以 0 为分割线，得到每个子区间的最大子集乘积
	lastStart := 0
	segNegSum := 0
	firstNegIdx, lastNegIdx := -1, -1

	for i := 0; i < len(nums); i++ {
		// 不好看的写法
		res = max(res, nums[i])

		if nums[i] < 0 {
			segNegSum++

			if firstNegIdx == -1 {
				firstNegIdx = i
			}
			lastNegIdx = i
		}
		if nums[i] == 0 {
			if lastStart != i {
				v := maxProductNoZero(nums[lastStart:i], segNegSum, firstNegIdx-lastStart, lastNegIdx-lastStart)
				res = max(res, v)
			}

			lastStart = i + 1
			segNegSum = 0
			firstNegIdx = -1
			lastNegIdx = -1
		}
	}

	v := res
	if lastStart < len(nums) && len(nums[lastStart:]) > 0 {
		v = maxProductNoZero(nums[lastStart:], segNegSum, firstNegIdx-lastStart, lastNegIdx-lastStart)
	}
	return max(res, v)
}

// 包含偶数个负数 → 全部相乘
// 包含奇数个负数 → 以第一个负数和最后一个负数为分界线，将区间划分成 3 段，取 max(第一段乘积 * 第二段乘积, 第二段乘积 * 第一段乘积)
func maxProductNoZero(nums []int, negSum, firstNegIdx, lastNegIdx int) int {
	// 一个数
	if len(nums) == 1 {
		return nums[0]
	}

	if negSum%2 == 0 {
		res := 1
		for i := 0; i < len(nums); i++ {
			res *= nums[i]
		}
		return res
	}

	one, two, three := 1, 1, 1
	for i := 0; i < firstNegIdx; i++ {
		one *= nums[i]
	}
	for i := firstNegIdx + 1; i < lastNegIdx; i++ {
		two *= nums[i]
	}
	for i := lastNegIdx + 1; i < len(nums); i++ {
		three *= nums[i]
	}

	// 只有一个负数
	if negSum == 1 {
		return max(one, three)
		// 超过一个负数
	} else {
		return max(one*nums[firstNegIdx]*two, two*nums[lastNegIdx]*three)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
