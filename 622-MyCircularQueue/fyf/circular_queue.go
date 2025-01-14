package fyf

type Node struct {
	value int
	next  *Node
}

type MyCircularQueue struct {
	head          *Node
	length        int
	currentLength int
	tail          *Node
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head:          &Node{},
		length:        k,
		currentLength: 0,
		tail:          &Node{},
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.currentLength >= this.length {
		return false
	}
	if this.currentLength == 0 {
		this.head = &Node{
			value: value,
			next:  nil,
		}
		this.currentLength++
		return true
	}

	h := this.head
	for h.next != nil {
		h = h.next
	}
	this.currentLength++
	if this.currentLength == this.length {
		h.next = &Node{
			value: value,
			next:  nil,
		}
		this.tail = &Node{
			value: value,
			next:  this.head,
		}
	} else {
		h.next = &Node{
			value: value,
			next:  nil,
		}
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.currentLength == 0 {
		return false
	}
	if this.currentLength == 1 {
		this.head = &Node{}
		this.currentLength--
		return true
	}
	if this.currentLength == this.length {
		this.tail.next = nil
	}

	this.head = this.head.next
	this.currentLength--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.currentLength == 0 {
		return -1
	}
	return this.head.value
}

func (this *MyCircularQueue) Rear() int {
	if this.currentLength == 0 {
		return -1
	}
	h := this.head
	for h.next != nil {
		h = h.next
	}
	return h.value
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.currentLength == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.currentLength == this.length {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
