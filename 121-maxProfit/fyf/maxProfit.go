package fyf

func maxProfit(prices []int) int {
	max := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		if max < prices[i]-minPrice {
			max = prices[i] - minPrice
		}
	}
	return max
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
