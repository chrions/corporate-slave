package fyf

import "math"

func maxSubArray(nums []int) int {
	minPreSum := 0
	sum := 0
	ret := math.MinInt
	for _, v := range nums {
		sum += v
		ret = max(ret, sum-minPreSum)
		minPreSum = min(sum, minPreSum)
	}

	return ret
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i <= j {
		return j
	}
	return i
}
