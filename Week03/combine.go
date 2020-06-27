package Week03

func combine(n int,k int) [][]int{
	if n<k{
		return [][]int{}
	}
	var res [][]int
	numSlice := make([]int,0)
	for i:=1;i<=k;i++{
		numSlice = append(numSlice,i)
	}
	dfsCombine(n,k,1,&res,[]int{})
	return res
}

func dfsCombine( n,k int, index int, res *[][]int, tmp []int) {
	if len(tmp) == k{
		*res = append(*res,tmp)
		return
	}
	for i:=index;i<=n;i++{
		tmpNow := dump(tmp)
		tmpNow = append(tmpNow,i)
		dfsCombine(n,k,i+1,res,tmpNow)
	}
}
func dump(a []int)[]int{
	b := make([]int,len(a))
	copy(b,a)
	return b
}

