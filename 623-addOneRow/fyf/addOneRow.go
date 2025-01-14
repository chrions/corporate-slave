package fyf

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}

	if depth == 1 {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}

	if depth == 2 {
		root.Left = &TreeNode{
			Val:   val,
			Left:  root.Left,
			Right: nil,
		}
		root.Right = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: root.Right,
		}
		return root
	}

	root.Left = addOneRow(root.Left, val, depth-1)
	root.Right = addOneRow(root.Right, val, depth-1)

	return root
}
