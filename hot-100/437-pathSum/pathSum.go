package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前缀和
func pathSum(root *TreeNode, targetSum int) int {
	// 使用哈希表记录前缀和的次数
	prefixSumCount := map[int]int{0: 1} // 0: 1 确保单节点路径被计算
	return dfs(root, 0, targetSum, prefixSumCount)
}

// 深度优先搜索（DFS）
func dfs(node *TreeNode, currSum, targetSum int, prefixSumCount map[int]int) int {
	if node == nil {
		return 0
	}

	// 更新当前前缀和
	currSum += node.Val

	// 计算满足条件的路径数
	count := prefixSumCount[currSum-targetSum]

	// 记录当前前缀和
	prefixSumCount[currSum]++

	// 递归遍历左右子树
	count += dfs(node.Left, currSum, targetSum, prefixSumCount)
	count += dfs(node.Right, currSum, targetSum, prefixSumCount)

	// **回溯（Backtracking）** - 恢复 `prefixSumCount` 状态，避免影响其他路径
	prefixSumCount[currSum]--

	return count
}

func pathSum2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	count := CountByNode(root, targetSum)
	count += pathSum2(root.Left, targetSum)
	count += pathSum2(root.Right, targetSum)
	return count
}

func CountByNode(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Val == targetSum {
		count++
	}
	count += CountByNode(node.Left, targetSum-node.Val)
	count += CountByNode(node.Right, targetSum-node.Val)
	return count
}
