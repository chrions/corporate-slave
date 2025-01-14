package fyf

import "sort"

func frequencySort(nums []int) []int {
	n := [][]int{}
	sort.Ints(nums)
	m := []int{}
	for i := 0; i < len(nums); i++ {
		m = append(m, nums[i])
		if i+1 < len(nums) && nums[i+1] != nums[i] {
			n = append(n, m)
			m = []int{}
		}
		if i == len(nums)-1 {
			n = append(n, m)
		}
	}
	sort.Slice(n, func(i, j int) bool {
		if len(n[i]) == len(n[j]) {
			return n[i][0] > n[j][0]
		}
		return len(n[i]) < len(n[j])
	})
	ret := make([]int, 0, len(nums))
	for _, v := range n {
		for _, vv := range v {
			ret = append(ret, vv)
		}

	}

	return ret
}
