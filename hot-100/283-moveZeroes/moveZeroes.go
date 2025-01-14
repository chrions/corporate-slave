package fyf

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if i+1 < len(nums) {
				for j := i + 1; j < len(nums); j++ {
					if nums[j] != 0 {
						nums[i], nums[j] = nums[j], nums[i]
						if j == len(nums)-1 {
							return
						}
						break
					}

				}
			}
		}
	}
}

func moveZeroes2(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
