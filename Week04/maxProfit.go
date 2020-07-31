package Week04

func maxProfit(prices []int) int {
	result := 0
	if len(prices) <= 1 {
		return 0
	}
	for idx := 1; idx < len(prices); idx++ {
		if prices[idx]-prices[idx-1] > 0 {
			result = result + (prices[idx] - prices[idx-1])
		}
	}
	return result
}
