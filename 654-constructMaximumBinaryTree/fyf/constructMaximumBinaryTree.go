package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	tree := buildTree(nums)
	return tree
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	rootIndex := 0
	root := 0
	for i, v := range nums {
		if v > root {
			root = v
			rootIndex = i
		}
	}
	tree := new(TreeNode)
	tree.Val = root

	leftNums := nums[0:rootIndex]
	rightNums := nums[rootIndex+1:]

	if rootIndex == 0 {
		leftNums = []int{}
	}
	if rootIndex == len(nums)-1 {
		rightNums = []int{}
	}
	tree.Left = buildTree(leftNums)
	tree.Right = buildTree(rightNums)
	return tree
}
