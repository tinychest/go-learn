package _0climbing_stairs

// Ques：You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Constraints:
//
//    1 <= n <= 45

func climbStairs(n int) int {
	// 超时了
	// return step1(n)

	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：1.9 MB, 在所有 Go 提交中击败了22.36% 的用户
	// 通过测试用例：45 / 45
	m := make(map[int]int, n)
	return step2(n, m)
}

func step1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 既然超时了，那就将中间计算数据存起来就行了
func step2(n int, m map[int]int) int {
	res, ok := m[n]
	if ok {
		return res
	}

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	res = step2(n-1, m) + step2(n-2, m)
	m[n] = res
	return res
}
