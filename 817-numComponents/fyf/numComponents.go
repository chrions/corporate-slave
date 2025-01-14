package fyf

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	ret := 0
	flag := false
	for head != nil {
		v := head.Val
		if _, ok := m[v]; ok {
			head = head.Next
			flag = true
			continue
		}
		if flag {
			ret++
		}
		flag = false
		head = head.Next
	}
	if flag {
		ret++
	}
	return ret
}
