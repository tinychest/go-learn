package _9combination_sum

// Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.
//
// The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.
//
// It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
//
//
// Constraints:
//
//    1 <= candidates.length <= 30
//    1 <= candidates[i] <= 200
//    All elements of candidates are distinct.
//    1 <= target <= 500

// - step1
// 就采用递归方式实现深度遍历的角度去做，但是，去重是一个问题
// 但是并不想通过笨拙的排序、去重来实现
// - step2
// 改变递归实现的思路，由题目意义的递归，变为采用 0~n 个目标数的递归
// - 其他
// 采用空间换时间的数值法，因为算法返回值（目标）是二维数组，所以实现起来比较复杂，所以跳过
func combinationSum(candidates []int, target int) [][]int {
	// return failStep1(candidates, nil, nil, target)
	return step2(candidates, nil, nil, target)
	// return step3(candidates, nil, nil, target)
}

func failStep1(candidates, selected []int, res [][]int, target int) [][]int {
	for _, v := range candidates {
		left := target - v
		if left == 0 {
			r := make([]int, len(selected)+1)
			copy(r, selected)
			r[len(selected)] = v
			res = append(res, r)
		} else if left > 0 {
			res = failStep1(candidates, append(selected, v), res, left)
		}
	}
	return res
}

// 提交是 fail 的，因为结果中每一个一维数组中的元素顺序是有要求的，虽然题目要求并没有写
func step2(candidates, selected []int, res [][]int, target int) [][]int {
	if len(candidates) == 0 {
		return res
	}

	for add := 0; add <= target; add += candidates[0] {
		if add == target {
			res = append(res, selected)
		}
		if add < target {
			res = step2(candidates[1:], selected, res, target-add)
		}
		selected = append(selected, candidates[0])
	}
	return res
}

// 只能转换思想，尽可能将指定元素用到最多，然后足部递减
func step3(candidates, selected []int, res [][]int, target int) [][]int {
	if len(candidates) == 0 {
		return res
	}
	value := candidates[0]
	times := target / value
	for i := 0; i < times; i++ {
		selected = append(selected, value)
	}
	for v := times * value; v >= 0; v -= value {
		if v == target {
			copied := make([]int, len(selected))
			copy(copied, selected)
			res = append(res, copied)
		}
		if v < target {
			res = step3(candidates[1:], selected, res, target-v)
		}

		if len(selected) != 0 {
			selected = selected[:len(selected)-1]
		}
	}
	return res
}
