package fyf

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := map[*TreeNode]struct{}{}
	exist := map[string]*TreeNode{}
	var help func(root *TreeNode) string
	help = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		serial := fmt.Sprintf("%d(%s)(%s)", root.Val, help(root.Right), help(root.Left))
		if n, ok := exist[serial]; ok {
			repeat[n] = struct{}{}
		} else {
			exist[serial] = root
		}
		return serial
	}
	help(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}
