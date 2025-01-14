package fyf

func getKthMagicNumber(k int) int {
	dp := make([]int, k+1)
	dp[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i <= k; i++ {
		v := min(dp[p3]*3, dp[p5]*5, dp[p7]*7)
		dp[i] = v
		if v == dp[p3]*3 {
			p3++
		}
		if v == dp[p5]*5 {
			p5++
		}
		if v == dp[p7]*7 {
			p7++
		}
	}
	return dp[k]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	if c <= a && c <= b {
		return c
	}
	return -1
}
