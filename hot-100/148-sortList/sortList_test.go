package fyf

import (
	"testing"
)

func TestSortList(t *testing.T) {
	list1 := &ListNode{Val: 4}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 1}
	list1.Next.Next.Next = &ListNode{Val: 3}
	ret := sortList(list1)
	t.Log(ret)
}
