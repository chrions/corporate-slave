package fyf

func groupThePeople(groupSizes []int) [][]int {
	m := map[int][]int{}
	for i := 0; i < len(groupSizes); i++ {
		m[groupSizes[i]] = append(m[groupSizes[i]], i)
	}
	ret := make([][]int, 0)
	for k, v := range m {
		if len(v) == k {
			ret = append(ret, v)
		}
		if len(v) > k {
			n := len(v) / k
			for i := 0; i < n; i++ {
				ret = append(ret, v[len(v)-k:])
				v = v[0 : len(v)-k]
			}
		}
	}
	return ret
}
