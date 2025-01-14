package fyf

type ListNode struct {
	Val  int
	Next *ListNode
}

// 数学推导 a = c + (n-1)(b+c)
// 也就是 a一定是n - 1圈加c的长度，这个时候，让ptr和慢指针同时+1着走，
// 那么慢指针走c之后，ptr到环入口的距离只剩整n - 1圈的距离了，而且此时慢指针也刚好走到入口处了
func detectCycle(head *ListNode) *ListNode {

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			temp := head
			if temp == slow {
				return slow
			}

			for temp != slow {
				temp = temp.Next
				slow = slow.Next
				if temp == slow {
					return slow
				}
			}
		}
	}
	return nil
}
