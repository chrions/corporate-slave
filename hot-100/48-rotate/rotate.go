package fyf

// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[7,4,1],[8,5,2],[9,6,3]]
func rotate(matrix [][]int) {
	n := len(matrix)
	temp := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		temp[i] = make([]int, n)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			temp[j][n-1-i] = matrix[i][j]
		}
	}
	copy(matrix, temp)
}
