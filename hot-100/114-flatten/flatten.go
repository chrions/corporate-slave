package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	ret := make([]*TreeNode, 0)
	stack := []*TreeNode{}
	for len(stack) != 0 || root != nil {
		for root != nil {
			ret = append(ret, root)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if node.Right != nil {
			root = node.Right
		}
	}
	for i := 1; i < len(ret); i++ {
		prev, curr := ret[i-1], ret[i]
		prev.Left, prev.Right = nil, curr
	}

	return
}

func flatten2(root *TreeNode) {

}
