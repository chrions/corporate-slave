package fyf

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {

			a := nums[i]
			b := nums[left]
			c := nums[right]
			if a+b+c == 0 {
				ret = append(ret, []int{a, nums[left], nums[right]})
				left++
				right--

				for right > left && nums[left] == b {
					left++
				}

				for right > left && nums[right] == c {
					right--
				}
			}
			if a+b+c > 0 {
				right--
			}
			if a+b+c < 0 {
				left++
			}
		}
	}
	return ret
}
