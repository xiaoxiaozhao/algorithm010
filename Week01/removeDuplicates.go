package Week01

//直接比较自己与自己相邻的数据的大小
//缺点，每次都要重新生成数组 效率较低
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	for idx := len(nums) - 1; idx > 0; idx-- {
		if nums[idx] == nums[idx-1] {
			nums = append(nums[:idx], nums[idx+1:]...)
		}
	}
	return len(nums)
}

//双指针法进行移动
func removeDuplicates_1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	left := 1
	right := 1
	for right < len(nums)-1 {
		if nums[right-1] != nums[right] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
