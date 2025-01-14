package fyf

import "testing"

func TestLongestUnivaluePath(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	ret := longestUnivaluePath(tree)
	t.Log(ret)
}
