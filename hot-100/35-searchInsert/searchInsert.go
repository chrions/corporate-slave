package fyf

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	temp := 0
	for start <= end {
		mid := (end + start) / 2
		temp = mid
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		}
		if nums[mid] < target {
			start = mid + 1
		}
	}
	if nums[temp] < target {
		return temp + 1
	}
	return temp
}
