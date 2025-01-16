package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var ret int

	if root == nil {
		return 0
	}
	exec(root, 0, &ret)

	return ret
}

func exec(node *TreeNode, deep int, ret *int) {
	if node == nil {
		return
	}
	deep++
	if deep > *ret {
		*ret = deep
	}
	exec(node.Left, deep, ret)
	exec(node.Right, deep, ret)

}

func maxDepth2(root *TreeNode) int {

	if root == nil {
		return 0
	}
	leftDeep := maxDepth2(root.Left)
	rightDeep := maxDepth2(root.Right)

	return max(leftDeep, rightDeep) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
