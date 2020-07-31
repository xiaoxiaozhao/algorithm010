package Week03

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootIndex := findIdx(inorder, preorder[0])
	if rootIndex == -1 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	//digui
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

func findIdx(nums []int, num int) int {
	for index, value := range nums {
		if value == num {
			return index
		}
	}
	return -1

}
