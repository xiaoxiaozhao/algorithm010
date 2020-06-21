package Week02

func nthUglyNumber(n int) int {
	dp := make([]int, 0)
	dp = append(dp, 1)
	n_two := 0
	n_three := 0
	n_five := 0
	for i := 1; i < n; i++ {
		tmp := min_1(dp[n_two]*2, min_1(dp[n_three]*3, dp[n_five]*5))
		dp = append(dp, tmp)
		if tmp == dp[n_two]*2 {
			n_two++
		}
		if tmp == dp[n_three]*3 {
			n_three++
		}
		if tmp == dp[n_five]*5 {
			n_five++
		}
	}
	return dp[n-1]
}

func min_1(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
