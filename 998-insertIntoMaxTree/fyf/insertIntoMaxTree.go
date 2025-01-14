package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if val > root.Val {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}
	if root.Right == nil {
		root.Right = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
		return root
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
