package fyf

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	nums := make([]int, len(nums1))
	ids := make([]int, len(nums2))
	for i := range ids {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return nums2[ids[i]] < nums2[ids[j]]
	})
	sort.Ints(nums1)
	left, right := 0, len(nums1)-1
	for _, v := range nums1 {
		if v > nums2[ids[left]] {
			nums[ids[left]] = v
			left++
		} else {
			nums[ids[right]] = v
			right--
		}
	}
	return nums
}
