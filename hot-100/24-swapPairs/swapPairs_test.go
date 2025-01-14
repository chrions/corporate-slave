package fyf

import "testing"

func TestSwapPairs(t *testing.T) {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 3}
	list1.Next.Next.Next = &ListNode{Val: 4}
	ret := swapPairs(list1)
	t.Log(ret)
}
