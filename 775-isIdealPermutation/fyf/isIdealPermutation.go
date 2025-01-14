package fyf

func isIdealPermutation(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i]-i > 1 {
			return false
		}
		if i+1 < len(nums) && nums[i] < nums[i+1] && nums[i] > i {
			return false
		}
	}
	return true
}
