package fyf

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}

	ret := removeNthFromEnd(list1, 2)
	t.Log(ret)
}
