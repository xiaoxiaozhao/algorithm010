package Week02

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	mem := make(map[int]int)
	for idx, num := range nums {
		other := target - num
		value, ok := mem[other]
		if ok {
			return []int{idx, value}
		} else {
			mem[num] = idx
		}
	}
	return []int{}
}
