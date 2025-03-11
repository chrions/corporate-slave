package fyf

type Point struct {
	i int
	j int
}

func orangesRotting(grid [][]int) int {
	directions := []Point{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	queue := []Point{}
	freshCount := 0

	for i, v := range grid {
		for j, _ := range v {
			if grid[i][j] == 2 {
				queue = append(queue, Point{i, j})
			}
			if grid[i][j] == 1 {
				freshCount++
			}
		}
	}
	if freshCount == 0 {
		return 0
	}

	minutes := -1 // 计时器（从 -1 开始，第一轮扩展后变 0）
	//BFS 进行腐烂扩散
	for len(queue) > 0 {
		nextQueue := []Point{} //下一轮要腐烂的
		for _, point := range queue {
			for _, direction := range directions {
				newI, newJ := point.i+direction.i, point.j+direction.j
				if isArea(grid, newI, newJ) && grid[newI][newJ] == 1 {
					grid[newI][newJ] = 2 //开始腐烂
					freshCount--         // 统计新鲜橘子减少
					nextQueue = append(nextQueue, Point{
						i: newI,
						j: newJ,
					})
				}
			}
		}
		queue = nextQueue
		minutes++
	}
	if freshCount > 0 {
		return -1
	}
	return minutes
}

func isArea(grid [][]int, i, j int) bool {
	if i < 0 || i > len(grid)-1 {
		return false
	}
	if j < 0 || j > len(grid[0])-1 {
		return false
	}
	return true
}
