package fyf

// 通过前缀和来计算 [i:j] 这个子数组和是否等于k
func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	m[0] = 1
	sum := 0
	ret := 0
	for _, v := range nums {
		sum += v
		if count, ok := m[sum-k]; ok {
			ret += count
		}
		m[sum]++
	}

	return ret
}
