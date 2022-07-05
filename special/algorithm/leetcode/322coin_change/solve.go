package _22coin_change

import (
	"fmt"
	"sort"
)

// [middle]
// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
//
// You may assume that you have an infinite number of each kind of coin.
//
// Constraints:
//
//    1 <= coins.length <= 12
//    1 <= coins[i] <= 231 - 1
//    0 <= amount <= 104
func coinChange(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	// return step1(coins, amount)
	// return step2(coins, amount)
	// return step3(coins, amount)
	return step4(coins, amount)
	// return refer(coins, amount)

	// 20 个数的答案测试

	// for i1 := 0; i1 <= 20; i1++ {
	// 	for i2 := 0; i2 <= 20; i2++ {
	// 		for i3 := 0; i3 <= 20; i3++ {
	// 			for i4 := 0; i4 <= 20; i4++ {
	// 				if i1 == 8 && i2 == 4 && i3 == 1 && i4 == 13 {
	// 					println()
	// 				}
	// 				if i1*419+i2*408+i3*186+i4*83 == 6249 {
	// 					fmt.Println(i1, i2, i3, i4)
	// 				}
	// 				// t.Log(8*419 + 4*408 + 1*186 + 13*83)
	// 			}
	// 		}
	// 	}
	// }
	// return 0
}

// 这是直觉的做法；很明显并不好，时间复杂度会很高，不然 leetcode 不会将其定级为 中等
func step1(coins []int, amount int) int {
	res := 0
	start := 0

	for i := 0; i < len(coins); i++ {
		if amount >= coins[i] {
			res++
			amount -= coins[i]

			if amount == 0 {
				return res
			}
			i = start - 1
		} else {
			start = i - 1
		}
	}
	if amount > 0 {
		return -1
	}
	return res
}

// 实际提交发现错误样例，就是本可以除尽，但因为使用了错误的硬币，导致除不尽了（step1 也存在一模一样的问题）
func step2(coins []int, amount int) int {
	res := 0

	for i := 0; i < len(coins); i++ {
		res += amount / coins[i]
		amount = amount % coins[i]

		if amount == 0 {
			return res
		}
	}
	if amount > 0 {
		return -1
	}
	return res
}

// 那只能进行扣减尝试，为了性能考虑，先逐渐减 1，而不是从 1 开始
// 实际证明这是错的，最少的数量似乎是没有规律的
func step3(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	coin := coins[0]

	di := amount / coin
	mo := amount % coin

	if mo == 0 {
		fmt.Printf("%d 个 %d\n", di, coin)
		return di
	}
	for j := di; j >= 0; j-- {
		tmp := step3(coins[1:], mo)
		if tmp != -1 {
			fmt.Printf("%d 个 %d\n", j, coin)
			return tmp + j
		}
		mo += coin
	}
	return -1
}

// 遍历的手段，提交 超时了...
func step4(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	// 达成目标数，所花的硬币数
	res := -1

	coin := coins[0]

	di := amount / coin
	mo := amount % coin

	if mo == 0 {
		return di
	}
	// 遍历所有情况，将最少的情况记录下来
	for j := di; j >= 0; j-- {
		tmp := step4(coins[1:], mo)
		if tmp != -1 {
			if res == -1 || res > tmp+j {
				res = tmp + j
			}
		}
		mo += coin
	}
	return res
}

// 其实 step4 已经非常接近这个答案了，step4 中仍然包含遍历的意思，而这里就像是通过借助额外的空间
// 达到了，略去判断给定的硬币能够凑到指定数的流程，所以是从小到大的过程
// 起始这道题和 70 Climbing Stairs 的原理非常相似
func refer(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = max
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
