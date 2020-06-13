package Week01

func rotate(nums []int, k int)  {
	if len(nums)<=1{
		return
	}
	//nums[0] = 100
	k = k%len(nums)
	result := append(nums[len(nums)-k:],nums[:len(nums)-k]...)
	for idx,item :=range result{
		nums[idx] = item
	}
}
