package fyf

import "testing"

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	got := mergeTwoLists(list1, list2)
	t.Log(got)

}
