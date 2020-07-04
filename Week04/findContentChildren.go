package Week04

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	//排序后进行一一对比
	count := 0
	gidx := 0
	for sidx:=0;sidx<=len(s)-1;sidx++{
		if gidx > len(g) -1 {
			break
		}
		if s[sidx] >= g[gidx]{
			count++
			gidx++
		}else {
			continue
		}
	}
	return count
}