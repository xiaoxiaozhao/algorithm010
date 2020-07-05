package Week04

import "fmt"

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0{
		return 0
	}
	column := len(grid[0])
	count := 0
	for i := 0 ;i < row; i++{
		for j := 0;j < column; j++ {
			//check if 1 or zero
			if grid[i][j] == '1'{
				count++
				processNums(grid,i,j)
				fmt.Print("1111")
			}

		}
	}
	return count
}
func processNums(nums [][]byte, i int, j int) {
	nums[i][j]= '0'
	if i-1 >= 0 {
		if nums[i-1][j] == '1' {
			nums[i-1][j] = '0'
			processNums(nums, i-1, j)
		}
	}
	if i+1 < len(nums) {
		if nums[i+1][j] == '1' {
			nums[i+1][j] = '0'
			processNums(nums, i+1, j)
		}
	}
	if j-1 >=0 {
		if nums[i][j-1] == '1' {
			nums[i][j-1] = '0'
			processNums(nums, i, j-1)
		}
	}
	if j+1 < len(nums[0]) {
		if nums[i][j+1] == '1' {
			nums[i][j+1] = '0'
			processNums(nums, i, j+1)
		}
	}


}
