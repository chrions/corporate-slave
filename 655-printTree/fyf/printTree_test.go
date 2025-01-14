package fyf

import "testing"

func TestPrintTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	ret := printTree(root)
	t.Log(ret)
}
