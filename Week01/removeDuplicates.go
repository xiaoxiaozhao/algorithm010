package Week01


func removeDuplicates(nums []int) int {
	if len(nums) <=1{
		return len(nums)
	}
	for idx:=len(nums)-1;idx>0;idx--{
		if nums[idx]==nums[idx-1]{
			nums = append(nums[:idx],nums[idx+1:]...)
		}
	}
	return len(nums)
}
