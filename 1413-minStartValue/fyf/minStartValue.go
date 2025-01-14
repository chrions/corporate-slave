package fyf

func minStartValue(nums []int) int {
	sum, min := 0, 0
	for _, v := range nums {
		sum += v
		min = getMin(sum, min)
	}
	return 1 - min
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
