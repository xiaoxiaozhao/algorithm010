package Week01

type Node struct {
	Next *Node
	Pre  *Node
	Val  int
}

type MyCircularDeque struct {
	Head        *Node
	Tail        *Node
	Capacity    int
	MaxCapacity int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	tail := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	head.Next = &tail
	tail.Pre = &head
	deque := MyCircularDeque{
		Head:        &head,
		Tail:        &tail,
		MaxCapacity: k,
		Capacity:    0,
	}

	return deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Capacity >= this.MaxCapacity {
		return false
	}
	tmp := this.Head.Next
	node := &Node{
		Pre:  this.Head,
		Next: this.Head.Next,
		Val:  value,
	}
	this.Head.Next = node
	tmp.Pre = node
	this.Capacity++
	return true

}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Capacity >= this.MaxCapacity {
		return false
	}
	tmp := this.Tail.Pre
	node := &Node{
		Pre:  this.Tail.Pre,
		Next: this.Tail,
		Val:  value,
	}
	this.Tail.Pre = node
	tmp.Next = node
	this.Capacity++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.Capacity == 0 {
		return false
	}
	itemDelete := this.Head.Next
	//删除开始
	this.Head.Next = itemDelete.Next
	itemDelete.Next.Pre = this.Head
	this.Capacity--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.Capacity == 0 {
		return false
	}
	itemDelete := this.Tail.Pre
	//删除开始
	this.Tail.Pre = itemDelete.Pre
	itemDelete.Pre.Next = this.Tail
	this.Capacity--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.Capacity == 0 {
		return this.Head.Val
	}
	return this.Head.Next.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.Capacity == 0 {
		return this.Tail.Val
	}
	return this.Tail.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.Capacity == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.Capacity == this.MaxCapacity {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
