package _22best_time_to_buy_and_sell_stock_II

// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
//
// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
// However, you can buy it then immediately sell it on the same day.
//
// Find and return the maximum profit you can achieve.
//
//
// Constraints:
//
//    1 <= prices.length <= 3 * 104
//    0 <= prices[i] <= 104

// 题目很好理解，相较于 121，不同的是，这里可以随时买进和卖出；太简单了，没有后续
func maxProfit(prices []int) int {
	var res, sub int
	for i := 0; i < len(prices)-1; i++ {
		sub = prices[i+1] - prices[i]
		if sub > 0 {
			res += sub
		}
	}
	return res
}
