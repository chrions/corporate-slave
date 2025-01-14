package fyf

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	h := &ListNode{Val: -1}
	h.Next = head
	pre := h
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return h.Next
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	n := make([]int, 0)
	n = append(n, head.Val)
	for head.Next != nil {
		head = head.Next
		n = append(n, head.Val)
	}
	n1 := n[left-1 : right]
	for i, j := 0, len(n1)-1; i < j; i, j = i+1, j-1 {
		n1[i], n1[j] = n1[j], n1[i]
	}
	n2 := append(append(n[0:left-1], n1...), n[right:]...)

	h := &ListNode{
		Val: n2[0],
	}
	temp := h
	for i := 1; i < len(n2); i++ {
		temp.Next = &ListNode{
			Val: n2[i],
		}
		temp = temp.Next
	}

	return h
}
