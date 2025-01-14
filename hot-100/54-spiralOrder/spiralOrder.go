package fyf

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix[0]), len(matrix)
	left, right, top, bottom := 0, rows-1, 0, columns-1
	ret := make([]int, 0, rows*columns)
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}

		for i := top + 1; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}

		if left < right && top < bottom {
			for i := right - 1; i > left; i-- {
				ret = append(ret, matrix[bottom][i])
			}
			for i := bottom; i > top; i-- {
				ret = append(ret, matrix[i][left])
			}
		}

		left++
		right--
		top++
		bottom--
	}
	return ret
}
