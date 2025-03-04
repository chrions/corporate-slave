package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 递归查找左右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果 p 和 q 分别在左右子树，则 root 是最近公共祖先
	if left != nil && right != nil {
		return root
	}

	// 否则返回非空的那个节点
	if left != nil {
		return left
	}
	return right
}
