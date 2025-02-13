package fyf

import "math"

// dp[i] = min(dp[i], dp[i-j*j]+1)
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minS := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minS = min(minS, dp[i-j*j]+1)
		}
		dp[i] = minS
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
