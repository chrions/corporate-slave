package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

// 前序 中左右
func preorderTraversal(root *TreeNode) []int {
	ret = make([]int, 0)
	getNode(root)
	return ret
}

func getNode(root *TreeNode) {
	if root == nil {
		return
	}
	ret = append(ret, root.Val)
	getNode(root.Left)
	getNode(root.Right)

}

// 迭代
func preorderTraversal1(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)

	current := root
	for len(stack) != 0 || current != nil {
		for current != nil {
			ret = append(ret, current.Val)
			stack = append(stack, current)
			current = current.Left
		}
		temp := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		current = temp.Right
	}
	return ret
}
