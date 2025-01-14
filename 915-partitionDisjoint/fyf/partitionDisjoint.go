package fyf

import "math"

func partitionDisjoint(nums []int) int {
	max := make([]int, len(nums))
	maxNums := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNums {
			max[i] = nums[i]
			maxNums = nums[i]
		} else {
			max[i] = maxNums
		}
	}

	min := make([]int, len(nums))
	minNums := math.MaxInt
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < minNums {
			min[i] = nums[i]
			minNums = nums[i]
		} else {
			min[i] = minNums
		}
	}
	for i, v := range min {
		if i > 0 && v >= max[i-1] {

			return i
		}
	}

	return 0
}
