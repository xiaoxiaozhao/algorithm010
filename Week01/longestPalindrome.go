package Week01

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	dp := make([][]int, n)
	for idx, _ := range dp {
		dp[idx] = make([]int, n)
	}
	left := 0
	right := 0
	for i := n - 2; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1] > 0) {
				dp[i][j] = 1
			}
			if dp[i][j] > 0 && j-i > right-left {
				right = j
				left = i
			}

		}
	}
	return s[left : right+1]
}

func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
