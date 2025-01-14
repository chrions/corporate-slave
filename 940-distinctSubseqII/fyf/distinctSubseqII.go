package fyf

func distinctSubseqII(s string) int {
	ans := 0
	const mod int = 1e9 + 7
	g := make([]int, 26)
	for _, c := range s {
		total := 1
		for _, v := range g {
			total = (total + v) % mod
		}
		g[c-'a'] = total
	}
	for _, v := range g {
		ans = (ans + v) % mod
	}
	return ans
}
