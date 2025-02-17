package fyf

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 1; i < n-1; i++ {
		row := make([]int, n-2)
		for j := 1; j < n-1; j++ {
			mx := grid[i][j]
			for r := i - 1; r <= i+1; r++ {
				for c := j - 1; c <= j+1; c++ {
					if grid[r][c] > mx {
						mx = grid[r][c]
					}
				}
			}
			row[j-1] = mx
		}
		ans[i-1] = row
	}
	return ans
}
