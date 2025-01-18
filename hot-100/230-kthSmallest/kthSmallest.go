package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	ret := make([]int, 0)
	stack := []*TreeNode{}
	for len(stack) != 0 || root != nil {

		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ret = append(ret, node.Val)
			if len(ret) == k {
				return ret[len(ret)-1]
			}
			root = node.Right
		}

	}
	return 0
}
