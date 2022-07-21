package _6permutations

// Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
//
//
// Constraints:
//
//    1 <= nums.length <= 6
//    -10 <= nums[i] <= 10
//    All the integers of nums are unique.

// [方向]
// 希望使用递归的方式去解决
// 尽可能复用空间的判断逻辑太复杂，就没打算复用空间了
//
// [结果]
// Runtime: 3 ms, faster than 57.23% of Go online submissions for Permutations.
// Memory Usage: 3.5 MB, less than 8.87% of Go online submissions for Permutations.
//
// [结论]
// 定义了不少辅助函数，实际是实现的空间利用效率太低了
//
// [参考]
// 又提到了回溯（backtracking），和 22 题是同性质的题
// 当然，这道题相较于括号结果的回溯，还需要额外借助一个符组的切片，代表当前已经被取用的元素
// 其实写这道题的时候就感觉，可以做到结果逐步累加，达到指定长度后添加到最终结果集中（复制是少不了的）
// `
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	res := make([][]int, 0, factorial(len(nums)))
	for i := 0; i < len(nums); i++ {
		for _, s := range permute(copyExcept(nums, i)) {
			s = append(copy(s), nums[i])
			res = append(res, s)
		}
	}
	return res
}

func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func copy(arr []int) []int {
	res := make([]int, 0, len(arr))
	for _, v := range arr {
		res = append(res, v)
	}
	return res
}

func copyExcept(arr []int, idx int) []int {
	res := make([]int, 0, len(arr)-1)
	for i, v := range arr {
		if i != idx {
			res = append(res, v)
		}
	}
	return res
}

// 参考的采用回溯的做法（为结果直接初始化好了最终的空间大小）
// func permute(nums []int) [][]int {
// 	permutations := make([][]int, 0, factorial(len(nums)))
// 	perm := make([]int, 0, len(nums))
//
// 	backtracking(&permutations, &perm, nums, make([]bool, len(nums)))
// 	return permutations
// }
//
// func backtracking(permutations *[][]int, perm *[]int, nums []int, chosen []bool) {
// 	if len(*perm) == len(nums) {
// 		newPerm := make([]int, len(*perm))
// 		for i := range *perm {
// 			newPerm[i] = (*perm)[i]
// 		}
// 		*permutations = append(*permutations, newPerm)
// 		return
// 	}
//
// 	for i := range nums {
// 		if chosen[i] {
// 			continue
// 		}
//
// 		*perm = append(*perm, nums[i])
// 		chosen[i] = true
//
// 		backtracking(permutations, perm, nums, chosen)
//
// 		*perm = (*perm)[:len(*perm)-1]
// 		chosen[i] = false
//
// 	}
// }