package fyf

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	cur := ret
	flag := 0
	for l1 != nil && l2 != nil {
		curValue := l1.Val + l2.Val + flag
		if curValue >= 10 {
			cur.Val = curValue - 10
			flag = 1
		} else {
			cur.Val = curValue
			flag = 0
		}
		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil || l2 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}

	restLink := func(l *ListNode) {
		for l != nil {
			curValue := l.Val + flag
			if curValue >= 10 {
				cur.Val = curValue - 10
				flag = 1
			} else {
				cur.Val = curValue
				flag = 0
			}
			l = l.Next
			if l != nil {
				cur.Next = new(ListNode)
				cur = cur.Next
			}
		}
		if flag == 1 {
			cur.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	}

	if l1 != nil {
		restLink(l1)
	} else if l2 != nil {
		restLink(l2)
	} else if flag == 1 {
		cur.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return ret
}

// addTwoNumbers2 gpt给出优化方案
// 在 for 循环中，我们一次性处理了两个链表的每一位加法。对于 l1 和 l2，如果当前节点不为空，则加上当前节点的值，并移动到下一个节点。加法和进位处理逻辑集中在循环内。
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个虚拟头节点，简化链表构建的逻辑
	ret := new(ListNode)
	cur := ret
	carry := 0

	// 遍历链表 l1 和 l2，直到两个链表都为空并且没有进位
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry // 先加上之前的进位

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next // 移动到下一个节点
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next // 移动到下一个节点
		}

		carry = sum / 10                    // 更新进位
		cur.Next = &ListNode{Val: sum % 10} // 保存当前位的结果
		cur = cur.Next                      // 移动当前指针
	}

	// 返回结果链表（跳过虚拟头节点）
	return ret.Next
}
