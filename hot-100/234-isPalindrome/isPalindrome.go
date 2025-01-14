package fyf

type ListNode struct {
	Val  int
	Next *ListNode
}

// 空间复杂度O(n)
func isPalindrome(head *ListNode) bool {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	i, j := 0, len(list)-1
	for i <= j {
		if list[i] != list[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 空间复杂度O(1)
func isPalindrome2(head *ListNode) bool {
	middleNode := middleListNode(head)
	reverseNode := reverseList(middleNode.Next)

	for reverseNode != nil {
		if head.Val != reverseNode.Val {
			return false
		}
		reverseNode = reverseNode.Next
		head = head.Next
	}

	return true
}

func middleListNode(list *ListNode) *ListNode {
	slow, fast := list, list
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}
