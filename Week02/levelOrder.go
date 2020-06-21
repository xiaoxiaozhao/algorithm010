package Week02

func levelOrder(root *Node) [][]int {
	if root == nil{
		return [][]int{}
	}
	queueMem := []*Node{root}
	var res [][]int
	for len(queueMem)!=0{
		tmplist := make([]int,0)
		for _,item := range queueMem{
			tmplist = append(tmplist,item.Val)
		}
		res = append(res,tmplist)
		tmpQueue := make([]*Node,0)
		for _,item :=range queueMem{
			for _,itemChild := range item.Children{
				tmpQueue = append(tmpQueue,itemChild)
			}
		}
		queueMem = tmpQueue
	}
	return res
}