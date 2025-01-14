package fyf

func productExceptSelf(nums []int) []int {
	l := make([]int, len(nums))
	r := make([]int, len(nums))
	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	r[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = l[i] * r[i]
	}
	return ret
}

func productExceptSelf2(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] = ret[i] * r
		r = r * nums[i]
	}
	return ret
}
