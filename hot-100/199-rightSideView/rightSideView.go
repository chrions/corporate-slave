package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		ret = append(ret, queue[len(queue)-1].Val)
		temp := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		queue = temp
	}
	return ret
}
