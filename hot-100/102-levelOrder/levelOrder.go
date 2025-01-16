package fyf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)

	ret := make([][]int, 0)

	queue = append(queue, root)
	for len(queue) != 0 {
		temp := make([]int, 0)
		queueTemp := make([]*TreeNode, 0)

		for i := 0; i < len(queue); i++ {
			temp = append(temp, queue[i].Val)

			if queue[i].Left != nil {
				queueTemp = append(queueTemp, queue[i].Left)
			}
			if queue[i].Right != nil {
				queueTemp = append(queueTemp, queue[i].Right)

			}
		}
		ret = append(ret, temp)
		queue = queueTemp
	}

	return ret
}
