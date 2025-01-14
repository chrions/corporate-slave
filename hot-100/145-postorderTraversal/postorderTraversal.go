package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

// 后序 左右中
func postorderTraversal(root *TreeNode) []int {
	ret = make([]int, 0)
	getNode(root)
	return ret
}

func getNode(root *TreeNode) {
	if root == nil {
		return
	}
	getNode(root.Left)
	getNode(root.Right)
	ret = append(ret, root.Val)

}

// 迭代
func postorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var pre *TreeNode

	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			//把左子树全部压栈
			stack = append(stack, current)
			current = current.Left
		}
		//取出栈顶元素
		top := stack[len(stack)-1]

		//如果右子树存在且未被访问，则处理右子树
		if top.Right != nil && top.Right != pre {
			current = top.Right
		} else {
			//否则，访问当前节点并从栈中弹出
			ret = append(ret, top.Val)
			stack = stack[0 : len(stack)-1]
			//标记当前节点为已访问
			pre = top
		}
	}
	return ret
}
