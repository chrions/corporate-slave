package fyf

import "testing"

// 左中右
func TestInorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	inorderTraversal(tree)
}
