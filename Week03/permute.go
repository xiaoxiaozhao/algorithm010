package Week03

func permute(nums []int) [][]int {
	length := len(nums)
	if length==0{
		return [][]int{}
	}
	var result [][]int
	searchPre(0,length,&result,[]int{},nums)
	return result
}

func searchPre(i int, length int, res *[][]int, sub []int,nums []int) {
	if i == length-1{
		sub = append(sub,nums[0])
		*res = append(*res,sub)
		return
	}
	for idx,value := range nums{
		//process
		tmp := make([]int,0)
		tmp = append(tmp,nums[0:idx]...)
		tmp = append(tmp,nums[idx+1:len(nums)]...)
		searchPre(i+1,length,res,append(sub,value),tmp)
	}
	return
}
