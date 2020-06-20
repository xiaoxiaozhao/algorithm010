package Week01

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	//var result []int
	mem := make(map[int]int)
	for idx, item := range nums {
		targetNum := target - item
		if idxTarget, ok := mem[targetNum]; ok {
			return []int{idx, idxTarget}
		} else {
			mem[item] = idx
		}
	}
	return []int{}
}
