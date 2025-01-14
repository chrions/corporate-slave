package fyf

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	nums1 = helper(nums1)
	nums2 = helper(nums2)
	nums3 = helper(nums3)

	m := map[int]int{}
	ret := make([]int, 0)
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		m[v]++
	}
	for _, v := range nums3 {
		m[v]++
	}
	for k, v := range m {
		if v > 1 {
			ret = append(ret, k)
		}
	}
	return ret
}

func helper(n []int) []int {
	m := map[int]struct{}{}
	ret := make([]int, 0)
	for _, v := range n {
		if _, ok := m[v]; ok {
			continue
		}
		ret = append(ret, v)
		m[v] = struct{}{}
	}
	return ret
}
