package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

// 中序 左中右
func inorderTraversal(root *TreeNode) []int {
	ret = make([]int, 0)
	getNode(root)
	return ret
}

func getNode(node *TreeNode) {
	if node == nil {
		return
	}
	getNode(node.Left)
	ret = append(ret, node.Val)
	getNode(node.Right)

}

// 迭代
func inorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			tempNode := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ret = append(ret, tempNode.Val)
			root = tempNode.Right
		}
	}
	return ret
}
