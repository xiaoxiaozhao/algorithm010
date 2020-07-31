package Week04

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	//查到所在row
	var numsTmp []int
	for idx := 0; idx < len(matrix); idx++ {
		numsTmp = append(numsTmp, matrix[idx][0])
	}
	//numsTmp := append([]int,...matrix[])
	rowIdx := midFindIdx2(numsTmp, target)
	if matrix[rowIdx][0] > target {
		rowIdx = rowIdx - 1
	}

	index := midFindIdx2(matrix[rowIdx], target)

	if matrix[rowIdx][index] == target {
		return true
	}
	return false
}

func midFindIdx2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	min := 0
	max := len(nums)
	mid := (min + max) / 2
	if nums[mid] == target {
		return mid
	}
	for {
		if min >= max || (max-min == 1) {
			break
		}
		if target > nums[mid] {
			min = mid
		} else if target < nums[mid] {
			max = mid
		} else {
			return mid
		}
		mid = (min + max) / 2
	}
	return mid

}
