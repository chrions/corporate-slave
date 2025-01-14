package fyf

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		m[sortStr([]byte(v))] = append(m[sortStr([]byte(v))], v)
	}
	ret := make([][]string, 0)
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func sortStr(s []byte) string {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return string(s)
}
