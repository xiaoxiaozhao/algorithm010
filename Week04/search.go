package Week04

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	for {
		if left >= right {
			if nums[left] == target {
				return left
			}
			break
		}
		if ((nums[0] > target) != (nums[0] > nums[mid])) != (target > nums[mid]) {
			left = mid + 1
			mid = (left + right) / 2
		} else {
			right = mid
			mid = (left + mid) / 2
		}
	}
	return -1
}
