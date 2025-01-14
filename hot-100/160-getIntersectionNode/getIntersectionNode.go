package fyf

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := make(map[*ListNode]struct{})
	for headA != nil {
		a[headA] = struct{}{}
		if headA.Next != nil {
			headA = headA.Next
			continue
		}
		break

	}
	for headB != nil {
		if _, ok := a[headB]; ok {
			return headB
		}
		if headB.Next != nil {
			headB = headB.Next
			continue
		}
		break
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	la, lb := headA, headB
	for la != lb {
		if la == nil {
			la = headB
		} else {
			la = la.Next
		}
		if lb == nil {
			lb = headA
		} else {
			lb = lb.Next
		}

	}
	return la
}
