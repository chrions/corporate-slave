package fyf

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用两个指针 slow 和 fast，初始时都指向虚拟头节点。通过让 fast 指针先走 n 步，
// 接着同时移动 slow 和 fast，直到 fast 到达链表的末尾。此时，slow 会指向倒数第 n 个节点的前一个节点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建一个虚拟头节点，方便处理删除头节点的情况
	dummy := &ListNode{Next: head}

	// 初始化快慢指针，指向虚拟头节点
	slow, fast := dummy, dummy

	// 让 fast 指针先走 n 步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 同时移动 slow 和 fast，直到 fast 到达链表的末尾
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 删除 slow.Next 节点（即倒数第 n 个节点）
	slow.Next = slow.Next.Next

	// 返回新的头结点（如果删除的是头结点，dummy.Next 会指向新的头结点）
	return dummy.Next
}
