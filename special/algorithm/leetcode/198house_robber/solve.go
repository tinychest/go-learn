package _98house_robber

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
//
// Constraints:
//
//    1 <= nums.length <= 100
//    0 <= nums[i] <= 400

// 思路1：1 2 3 4 ...，抢劫 1 能否取得最大价值，取决于 2 是否大于 1 + 3，然后，看 3，之后其实就是同理，要不要抢 3 取决于 3 + 5 是否大于 4
//      发现有问题，实际应该是，如果一个房子的价值是否大于相邻两个房子的价值之和，那么这个房子是一定要选的，反之，并不对
// 思路2：顺着上面的思路，并不能得到有效的答案，细想一组元素，无论价值是如何分布，第一个、第二个元素肯定是要选择一个的
//      但是，这并不能得出，两位两位考虑即可的结论
// 参考：这种题是动态规化题，应该在过程中逐步得到最终的结果，并不能以找到题目特征规律的方式得到最终答案
func rob(nums []int) int {
	// return try1(nums)
	// return try2(nums)
	return final(nums)
}

func final(nums []int) int {
	cur, pre, ppre := 0, 0, 0
	for _, v := range nums {
		cur = max(pre, ppre+v)
		ppre, pre = pre, cur
	}
	return cur
}

// 很明显和汉诺塔一样，这里重复计算了太多中间已经得出的答案，可以使用 dp 存储中间结果，来提升效率
// （自己的角度是这样，实际上 DP 数组应该称之为 子问题数组，因为数组中的每一个元素都对应一个子问题）
//
// 这道题的一个题友解答拓展的很广：
// - 这道题可以作为一道 动态规化 的入门题
// - 使用 dp 数组可以轻易的简化为使用两个变量即可
func try2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}
	return max(try2(nums[:len(nums)-1]), try2(nums[:len(nums)-2])+nums[len(nums)-1])
}

func try1(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if i == len(nums)-1 {
			res += nums[i]
			break
		} else if i == len(nums)-2 {
			res += max(nums[i], nums[i+1])
			break
		} else {
			if nums[i+1] > nums[i]+nums[i+2] {
				res += nums[i+1]
				i = i + 1 + 1 // 加的数的下标 + 往后跳 1 个 + 循环的自增，相当于直接跳过相邻的数
			} else {
				res += nums[i]
				i = i + 1
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
