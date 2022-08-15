package _23best_time_to_buy_and_sell_stock_III

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// Find the maximum profit you can achieve. You may complete at most two transactions.
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
//
// Constraints:
//
//    1 <= prices.length <= 105
//    0 <= prices[i] <= 105

// 题目很好理解，相较于 122，这里只能交易两笔，也就是说取两端涨的最多的就行，赚最大的两笔就行
func maxProfit(prices []int) int {
	changes := make([]int, len(prices)-1)
	v, sub := 0, 0

	for i := 0; i < len(prices)-1; i++ {
		sub = prices[i+1] - prices[i]

		if (v > 0 && sub < 0) || (v < 0 && sub > 0) {
			changes = append(changes, v)
			v = sub
		} else {
			v += sub
		}
	}

	// 从 changes 中划取两段连续的区间，使得区间中的值的和最大
	// - 到了这里，没那么简单，需要尝试进一步合并区间，合并到最终结果
	//   （需要先思考好，如何决策分组，能使得结果最大）
	//   如果合并的结果区间等于 1，那只要减去亏的最多的区间的值之和就是答案
	//   大于 1，就是将区间的值之和进行排序，选两个最大的
	// - 分析演进到上面那样，还是抽象，想办法遍历出最佳解

	return 0
}
