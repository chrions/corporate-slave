package fyf

type Node struct {
	value int
	next  *Node
	pre   *Node
}

type MyCircularDeque struct {
	length        int
	currentLength int
	head          *Node
	tail          *Node
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		length:        k,
		currentLength: 0,
		head:          nil,
		tail:          nil,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.currentLength == 0 {
		n := &Node{
			value: value,
			next:  nil,
			pre:   nil,
		}
		this.head = n
		this.tail = n
		this.currentLength++
		return true
	}
	if this.currentLength < this.length {
		h := &Node{
			value: value,
			next:  nil,
			pre:   nil,
		}

		this.head.pre = h
		h.next = this.head
		this.head = h
		this.currentLength++
		return true
	}
	return false
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.currentLength == 0 {
		n := &Node{
			value: value,
			next:  nil,
			pre:   nil,
		}
		this.head = n
		this.tail = n
		this.currentLength++
		return true
	}
	if this.currentLength < this.length {
		t := &Node{
			value: value,
			next:  nil,
			pre:   nil,
		}
		this.tail.next = t
		t.pre = this.tail
		this.tail = t
		this.currentLength++
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.currentLength == 1 {
		this.head = nil
		this.tail = nil
		this.currentLength--
		return true
	}
	if this.currentLength == 2 {
		n := this.head.next
		n.pre = nil
		n.next = nil
		this.head = n
		this.tail = n

		this.currentLength--
		return true
	}
	if this.currentLength > 2 {
		h := this.head.next
		this.head = h
		h.pre = nil
		this.currentLength--
		return true
	}

	return false
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.currentLength == 1 {
		this.head = nil
		this.tail = nil
		this.currentLength--
		return true
	}
	if this.currentLength == 2 {
		n := this.tail.pre
		n.pre = nil
		n.next = nil
		this.head = n
		this.tail = n
		this.currentLength--
		return true
	}
	if this.currentLength > 2 {
		n := this.tail.pre
		this.tail = n
		n.next = nil
		this.currentLength--
		return true
	}
	return false
}

func (this *MyCircularDeque) GetFront() int {
	if this.currentLength == 0 {
		return -1
	}
	return this.head.value
}

func (this *MyCircularDeque) GetRear() int {
	if this.currentLength == 0 {
		return -1
	}
	return this.tail.value
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.currentLength == 0 {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if this.currentLength == this.length {
		return true
	}
	return false
}
