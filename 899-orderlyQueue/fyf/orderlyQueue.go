package fyf

import "sort"

var minS string

func orderlyQueue(s string, k int) string {
	if len(s) == 1 {
		return s
	}
	minS = s
	if k == 1 {
		newS := ""
		ss := s
		m := map[string]struct{}{}
		for newS != ss {
			newS = s[1:len(s)] + string(s[0])
			m[newS] = struct{}{}
			s = newS
		}
		for k, _ := range m {
			if minS > k {
				minS = k
			}
		}
	}
	if k > 1 {
		b := make([]byte, 0, k)
		for i := 0; i < len(s); i++ {
			b = append(b, s[i])
		}
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		ss := ""
		for _, v := range b {
			ss += string(v)
		}
		minS = ss
	}
	return minS
}
