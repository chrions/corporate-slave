package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	var deep func(node *TreeNode) int
	deep = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftLength := deep(node.Left)
		rightLength := deep(node.Right)
		ret = max(ret, leftLength+rightLength)
		return max(leftLength, rightLength) + 1
	}
	deep(root)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
