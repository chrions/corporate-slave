package fyf

func setZeroes(matrix [][]int) {
	zeroX := map[int]struct{}{}
	zeroY := map[int]struct{}{}
	for i, v := range matrix {
		for j, vv := range v {
			if vv == 0 {
				zeroX[j] = struct{}{}
				zeroY[i] = struct{}{}
			}
		}
	}

	for i, v := range matrix {
		if _, ok := zeroY[i]; ok {
			matrix[i] = make([]int, len(v))
			continue
		}
		for j, _ := range v {
			if _, ok := zeroX[j]; ok {
				v[j] = 0
			}
		}
	}
}
