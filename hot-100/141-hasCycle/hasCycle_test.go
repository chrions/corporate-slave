package fyf

import "testing"

func TestHasCycle(t *testing.T) {

	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	ret := hasCycle(node)
	t.Log(ret)
}
