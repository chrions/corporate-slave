package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var symmetric bool

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	symmetric = true
	exec(root.Left, root.Right)
	return symmetric
}

func exec(left, right *TreeNode) {
	if left == nil && right != nil {
		symmetric = false
		return
	}
	if left != nil && right == nil {
		symmetric = false
		return
	}
	if left == nil && right == nil {
		return
	}

	if left.Val != right.Val {
		symmetric = false
		return
	}
	exec(left.Left, right.Right)
	exec(left.Right, right.Left)
}

// 迭代
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return false
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root.Left, root.Right)
	for len(stack) != 0 {
		left, right := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[0 : len(stack)-2]

		if left == nil && right == nil {
			continue
		}

		// 一个为空一个不为空，或者值不同，不对称
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		stack = append(stack, left.Left, right.Right)
		stack = append(stack, left.Right, right.Left)

	}
	return true
}
