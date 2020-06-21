package Week02

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int,0)
	if root == nil{
		return res
	}
	stackMem := make([]*TreeNode,0)
	stackMem = append(stackMem,root)
	for len(stackMem) != 0 {
		tmp := stackMem[len(stackMem)-1]
		if tmp.Left != nil{
			stackMem = append(stackMem,tmp.Left)
		}else {
			res = append(res,tmp.Val)
			stackMem = stackMem[0 : len(stackMem)-1]
			if len(stackMem) !=0 {
				stackMem[len(stackMem)-1].Left = nil
			}
			if tmp.Right != nil {
				stackMem = append(stackMem, tmp.Right)
			}
		}
	}
	return res
}