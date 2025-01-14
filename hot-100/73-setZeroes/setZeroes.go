package fyf

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix[0]))
	lines := make([]bool, len(matrix))
	for y, v := range matrix {
		for x, vv := range v {
			if vv == 0 {
				rows[x] = true
				lines[y] = true
			}
		}
	}
	for y, v := range matrix {
		for x, _ := range v {
			if rows[x] || lines[y] {
				matrix[y][x] = 0
			}
		}
	}
}
