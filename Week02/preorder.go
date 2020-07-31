package Week02

type Node struct {
	Val      int
	Children []*Node
}

//stack
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stackMem := make([]*Node, 0)
	//append root
	stackMem = append(stackMem, root)
	for len(stackMem) != 0 {
		item := stackMem[len(stackMem)-1]
		result = append(result, item.Val)
		stackMem = stackMem[0 : len(stackMem)-1]
		for i := 0; i <= len(item.Children)-1; i++ {
			if item.Children[len(item.Children)-i-1] != nil {
				stackMem = append(stackMem, item.Children[len(item.Children)-i-1])
			}
		}
	}
	return result
}

//circle
func preorder1(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	result = append(result, root.Val)
	//result = append(result,root.Val)
	for _, item := range root.Children {
		//	result = append(result,root.Val)
		result = append(result, preorder1(item)...)
	}
	return result
}
