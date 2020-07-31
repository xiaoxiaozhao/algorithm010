package Week04

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	rightMax := 0
	for idx, value := range nums {
		if idx <= rightMax {
			rightMax = max(rightMax, value+idx)
			if rightMax >= len(nums)-1 {
				return true
			}
		}
	}
	return false

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
