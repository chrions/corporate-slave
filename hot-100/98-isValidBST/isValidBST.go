package fyf

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	preOrder := math.MinInt
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left

		} else {
			node := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if node.Val <= preOrder { //bst的中序遍历是递增的，所以当前节点的值应该大于上一个访问的节点值。
				return false
			}
			preOrder = node.Val
			root = node.Right
		}
	}
	return true
}
