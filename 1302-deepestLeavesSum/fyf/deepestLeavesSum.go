package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//广度优先
func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		tempNode := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Left != nil {
				tempNode = append(tempNode, node.Left)
			}
			if node.Right != nil {
				tempNode = append(tempNode, node.Right)
			}
		}
		if len(tempNode) == 0 {
			break
		}
		queue = tempNode

	}
	sum := 0
	for _, v := range queue {
		sum += v.Val
	}
	return sum
}

//深度优先
func deepestLeavesSum2(root *TreeNode) int {
	maxDepp := 0
	ret := 0
	var dfs func(root *TreeNode, deep int)
	dfs = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}

		if deep > maxDepp {
			maxDepp = deep
			ret = root.Val
		} else if deep == maxDepp {
			ret += root.Val
		}

		dfs(root.Left, deep+1)
		dfs(root.Right, deep+1)
	}
	dfs(root, 1)
	return ret
}
