package fyf

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	start := getStart(nums, target)
	end := getEnd(nums, target)
	ret[0] = start
	ret[1] = end
	return ret
}

func getStart(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		}
		if nums[mid] < target {
			start = mid + 1
		}
		if nums[mid] > target {
			end = mid - 1
		}
	}
	return -1
}

func getEnd(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			start = mid + 1
		}
		if nums[mid] < target {
			start = mid + 1
		}
		if nums[mid] > target {
			end = mid - 1
		}
	}
	return -1
}
