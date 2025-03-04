package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	tree := new(TreeNode)
	tree.Val = preorder[0]

	root := preorder[0]
	index := 0
	for i, v := range inorder {
		if v == root {
			index = i
			break
		}
	}
	tree.Left = buildTree(preorder[1:index+1], inorder[0:index])
	tree.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return tree
}
