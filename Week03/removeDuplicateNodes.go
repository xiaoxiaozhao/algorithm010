package Week03

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := head
	mem := make(map[int]bool)
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	for head != nil {
		if _, ok := mem[head.Val]; ok {
			pre.Next = head.Next
		} else {
			//add
			pre = pre.Next
			mem[head.Val] = true
		}
		//next
		head = head.Next
	}
	return dummy.Next
}
