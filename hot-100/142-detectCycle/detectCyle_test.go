package fyf

import "testing"

func TestDetectCycle(t *testing.T) {

	ode := &ListNode{
		Val:  1,
		Next: nil,
	}
	ode2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	ode.Next = ode2
	ode2.Next = ode

	ret := detectCycle(ode)
	t.Log(ret)
}
