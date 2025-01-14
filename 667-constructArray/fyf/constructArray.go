package fyf

func constructArray(n int, k int) []int {
	ret := make([]int, n)

	m := map[int]struct{}{}
	for i := 0; i < n; i++ {
		if i == 0 {
			ret[i] = 1
			m[ret[i]] = struct{}{}
			continue
		}
		if i == 1 {
			ret[i] = 1 + k
			m[ret[i]] = struct{}{}
			continue
		}

		if k == 1 {
			ret[i] = ret[i-1] + 1
			continue
		}
		if abs(ret[i-2], ret[i-1])-1 == 0 {
			k = 1
			ret[i] = ret[1] + 1
			continue
		}
		v := ret[i-1] - (abs(ret[i-2], ret[i-1]) - 1)
		if v <= 0 {
			v = ret[i-1] + (abs(ret[i-2], ret[i-1]) - 1)
		}
		if _, ok := m[v]; ok {
			ret[i] = ret[i-1] + (abs(ret[i-2], ret[i-1]) - 1)
			continue
		}
		m[v] = struct{}{}
		ret[i] = v
	}
	return ret
}

func abs(a, b int) int {
	n := a - b
	if n >= 0 {
		return n
	}
	return -n
}
