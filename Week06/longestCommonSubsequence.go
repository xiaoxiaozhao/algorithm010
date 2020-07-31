package Week06

import "math"

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	//dp start
	str1Len := len(text1)
	str2Len := len(text2)
	res := make([][]int, str1Len)
	for i := 0; i < str1Len; i++ {
		res[i] = make([]int, str2Len)
	}
	//初始化边界
	maxValue := math.MinInt32
	for i := 0; i < str1Len; i++ {
		if text1[i] == text2[0] {
			res[i][0] = 1
			maxValue = 1
		} else {
			res[i][0] = max(maxValue, 0)
		}
	}
	maxValue = math.MinInt32
	for i := 0; i < str2Len; i++ {
		if text2[i] == text1[0] {
			res[0][i] = 1
			maxValue = 1
		} else {
			res[0][i] = max(maxValue, 0)
		}
	}

	for i := 1; i < str1Len; i++ {
		for j := 1; j < str2Len; j++ {
			//dp
			if text1[i] == text2[j] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = max(res[i][j-1], res[i-1][j])
			}
		}
	}
	return res[str1Len-1][str2Len-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
