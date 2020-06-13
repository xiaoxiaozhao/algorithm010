package Week01
type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	headPoint := &head
	tmpl1 := l1
	tmpl2 := l2
	for{
		tmp := new(ListNode)
		if tmpl1 == nil && tmpl2 == nil{
			break
		}
		if tmpl1 == nil{
			tmp = tmpl2
			headPoint.Next = tmp
			break
		}else if tmpl2 == nil{
			tmp = tmpl1
			headPoint.Next = tmp
			break
		}else {
			if tmpl1.Val <= tmpl2.Val{
				tmp.Val = tmpl1.Val
				tmpl1 = tmpl1.Next
			}else {
				tmp.Val = tmpl2.Val
				tmpl2 = tmpl2.Next
			}
			headPoint.Next = tmp
			headPoint = headPoint.Next
		}

	}
	return head.Next
}