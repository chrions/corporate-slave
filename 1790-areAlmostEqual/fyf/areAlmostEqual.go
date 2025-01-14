package fyf

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	m := make([]int, 0)
	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			m = append(m, i)
		}
	}
	if len(m) != 2 {
		return false
	}

	if s1[m[0]] == s2[m[1]] && s1[m[1]] == s2[m[0]] {
		return true
	}
	return false
}
