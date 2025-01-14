package fyf

import "sort"

func minSubsequence(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	ret := make([]int, 0)
	sort.Ints(nums)

	currentCount := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if currentCount > getNumsTotal(nums[0:i+1]) {
			ret = append(ret, nums[i+1:]...)
			break
		}
		currentCount += nums[i]
	}
	if len(ret) == 0 {
		return nums
	}

	//反转数组
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

func getNumsTotal(num []int) int {
	count := 0
	for _, v := range num {
		count += v
	}
	return count
}
