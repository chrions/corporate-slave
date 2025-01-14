package fyf

import "sort"

func minOperations(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	sum1 := 0
	for i := 0; i < len(nums1); i++ {
		sum1 += nums1[i]
	}
	sum2 := 0
	for i := 0; i < len(nums2); i++ {
		sum2 += nums2[i]
	}
	if sum2 == sum1 {
		return 0
	}
	ret := 0
	if sum2 > sum1 {
		diff := sum2 - sum1
		ret = helper(diff, nums1, nums2)

	}
	if sum2 < sum1 {
		diff := sum1 - sum2
		ret = helper(diff, nums2, nums1)
	}

	return ret
}

func helper(diff int, nums1, nums2 []int) int {
	ret := 0
	i := 0
	j := len(nums2) - 1
	for i <= len(nums1)-1 || j >= 0 {
		n := 0
		m := 0
		if i <= len(nums1)-1 {
			n = 6 - nums1[i]
		}
		if j >= 0 {
			m = nums2[j] - 1
		}
		if n == 0 && m == 0 {
			break
		}

		if n <= m {
			diff = diff - (nums2[j] - 1)

			ret++
			j--
		}
		if n > m {
			diff = diff - (6 - nums1[i])

			ret++
			i++
		}
		if diff <= 0 {
			return ret
		}
	}
	if diff <= 0 {
		return ret
	}
	return -1
}
