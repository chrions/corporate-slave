package fyf

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	height := geHeight(root) - 1
	res := make([][]string, height+1)
	n := int(math.Exp2(float64(height+1)) - 1)
	for i := range res {
		res[i] = make([]string, n)
	}
	cow := 0
	column := (n - 1) / 2
	fillRes(root, res, cow, column, height)

	return res
}

func fillRes(root *TreeNode, res [][]string, cow int, column, height int) {
	if root == nil {
		return
	}
	res[cow][column] = fmt.Sprintf("%d", root.Val)
	fillRes(root.Left, res, cow+1, column-int(math.Exp2(float64(height-cow-1))), height)
	fillRes(root.Right, res, cow+1, column+int(math.Exp2(float64(height-cow-1))), height)

}

func geHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(geHeight(root.Left), geHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
