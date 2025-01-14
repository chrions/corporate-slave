package fyf

func finalPrices(prices []int) []int {
	ret := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		ret[i] = prices[i]

		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				ret[i] = prices[i] - prices[j]
				break
			}
		}
	}
	ret[len(ret)-1] = prices[len(prices)-1]
	return ret
}
