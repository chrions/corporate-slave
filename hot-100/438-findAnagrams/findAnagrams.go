package fyf

func findAnagrams(s string, p string) []int {
	ret := make([]int, 0)
	if len(p) > len(s) {
		return ret
	}
	countP := [26]int{}
	countS := [26]int{}
	for i := 0; i < len(p); i++ {
		countP[p[i]-'a']++
		countS[s[i]-'a']++
	}
	if countP == countS {
		ret = append(ret, 0)
	}
	length := len(p)
	j := 0
	for i := length; i < len(s); i++ {
		countS[s[i]-'a']++
		countS[s[j]-'a']--
		j++
		if countP == countS {
			ret = append(ret, j)
		}
	}
	return ret
}
