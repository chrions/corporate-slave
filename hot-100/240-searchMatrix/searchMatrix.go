package fyf

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	// 从右上角开始搜索
	row := 0
	col := n - 1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col-- // 当前值大于目标值,向左移动
		} else {
			row++ // 当前值小于目标值,向下移动
		}
	}
	return false
}
