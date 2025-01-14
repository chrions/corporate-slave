package fyf

func shuffle(nums []int, n int) []int {
	ret := make([]int, 0, len(nums))
	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		ret = append(ret, nums[i], nums[j])
	}
	return ret
}
