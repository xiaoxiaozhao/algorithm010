package Week04

func findFirstNotSortIdx(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mid := len(nums) / 2
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[mid] > nums[mid + 1] {
			return nums[mid + 1]
		}

		// if the mid element is lesser than its previous element then mid element is the smallest
		if nums[mid - 1] > nums[mid] {
			return nums[mid]
		}

		//if nums[mid-1] < nums[mid-2] && nums[mid-1] < nums[mid-1]{
		//	return mid
		//}
		if nums[left] > nums[mid] {
			right = mid
		} else {
			left = mid
		}
		mid = (left + right) / 2
	}
	return 0
}
