package Week03

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0{
		return [][]int{}
	}
	var res [][]int
	searchPer(nums,0,&res)
	return res
}

func searchPer(nums []int, index int, res *[][]int) {
	if index == len(nums){
		*res = append(*res,dump(nums))
		return
	}
	mem := make(map[int]bool,0)
	for k:=index;k<len(nums);k++{
		if _,ok:=mem[nums[k]];ok{
			continue
		}
		nums[index],nums[k] = nums[k],nums[index]
		searchPer(nums,index+1,res)
		nums[index],nums[k] = nums[k],nums[index]
		mem[nums[k]] = true
	}
}
