package fyf

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return mergeSort(head)
}

func mergeSort(l *ListNode) *ListNode {
	if l.Next == nil {
		return l
	}
	fast, slow := l, l
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil

	return merge(mergeSort(l), mergeSort(mid))
}

func merge(l1, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	cur := ret
	for l1 != nil && l2 != nil {
		node := new(ListNode)
		if l1.Val < l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}
		cur.Next = node
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return ret.Next
}
