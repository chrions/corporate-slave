package fyf

func numIslands(grid [][]byte) int {
	ret := 0
	for i, v := range grid {
		for j, _ := range v {
			if string(grid[i][j]) == "1" {
				ret++
				dsf(grid, i, j)
			}
		}
	}
	return ret
}

func dsf(grid [][]byte, i, j int) {
	if !IsArea(grid, i, j) {
		return
	}

	if string(grid[i][j]) != "1" {
		return
	}

	grid[i][j] = byte(2)

	dsf(grid, i+1, j)
	dsf(grid, i-1, j)
	dsf(grid, i, j-1)
	dsf(grid, i, j+1)

}

func IsArea(grid [][]byte, i, j int) bool {
	if i > len(grid)-1 || i < 0 {
		return false
	}
	if j > len(grid[0])-1 || j < 0 {
		return false
	}
	return true
}
