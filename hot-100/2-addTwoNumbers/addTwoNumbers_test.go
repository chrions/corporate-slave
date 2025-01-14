package fyf

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	list1 := &ListNode{Val: 5}

	list2 := &ListNode{Val: 5}

	ret := addTwoNumbers(list1, list2)
	t.Log(ret)
}
