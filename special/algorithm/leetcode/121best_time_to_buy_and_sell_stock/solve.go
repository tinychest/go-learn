package _21best_time_to_buy_and_sell_stock

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
//
// Constraints:
//
//    1 <= prices.length <= 105
//    0 <= prices[i] <= 104

func maxProfit(prices []int) int {
	// return step1(prices)
	return step2(prices)
}

// 核心思想就和两块板最多能容纳多少水那道题一样，但是提交结果并不理想
// 这就不用纠结了，事件复杂度就是 O(n) 已经是很不错的解法了
func step2(prices []int) int {
	var res int
	a, b := prices[0], prices[0]
	for _, v := range prices {
		if v < a {
			a = v
			b = v
		}
		if v > b {
			b = v
		}
		if res < b-a {
			res = b - a
		}
	}
	return res
}

func step1(prices []int) int {
	// 在指定日期买入，并向后找一个能达到的最大值
	// 先用死办法实现一下
	var res int
	for i := 0; i < len(prices)-1; i++ {
		max := prices[i]
		for j := i + 1; j < len(prices)-1; j++ {
			if max < prices[j] {
				max = prices[j]
			}
		}
		if res < max-prices[i] {
			res = max - prices[i]
		}
	}
	return res
}
