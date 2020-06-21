package Week02
func preorderTraversal(root *TreeNode) []int {
	res := make([]int,0)
	if root == nil{
		return res
	}
	stackMem := make([]*TreeNode,0)
	stackMem = append(stackMem,root)
	for len(stackMem) != 0 {
		//先进先出
		res = append(res,stackMem[len(stackMem)-1].Val)
		//stackMem = stackMem[0:len(stackMem)-1]
		tmp := stackMem[len(stackMem)-1]
		stackMem = stackMem[0:len(stackMem)-1]
		if tmp.Right != nil {
			stackMem = append(stackMem, tmp.Right)
		}
		if tmp.Left != nil{
			stackMem = append(stackMem, tmp.Left)
		}
	}
	return res
}