package leetcode

import "math"

func maxProfit(prices []int) int {
	// 问题其实为 取一个数左边，一个数右边，最大差值
	// 找最小值一定是必须且固定的
	minPrice := math.MaxInt64
	res := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		profit := prices[i] - minPrice
		if profit > res {
			res = profit
		}
	}
	return res
}
