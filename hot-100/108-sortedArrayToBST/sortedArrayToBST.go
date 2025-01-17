package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return initTree(nums, 0, len(nums)-1)
}

func initTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (right + left + 1) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = initTree(nums, left, mid-1)
	root.Right = initTree(nums, mid+1, right)

	return root
}
