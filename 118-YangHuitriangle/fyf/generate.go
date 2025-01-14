package fyf

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		n := make([]int, i)
		n[0] = 1
		n[i-1] = 1
		if i > 2 {
			m := ret[i-2]
			for j := 0; j < len(m); j++ {
				if j+1 < len(m) {
					v := m[j] + m[j+1]
					n[j+1] = v
				}

			}
		}
		ret[i-1] = n
	}
	return ret
}
