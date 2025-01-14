package fyf

func rotate(nums []int, k int) {
	rest := k % len(nums)
	if rest == 0 {
		return
	}
	nums = append(nums[len(nums)-rest:], nums[0:len(nums)-rest]...)
}
