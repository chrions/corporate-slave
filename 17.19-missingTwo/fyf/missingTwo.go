package fyf

func missingTwo(nums []int) []int {
	item := make([]int, len(nums)+2)
	for _, v := range nums {
		item[v-1] = v
	}
	flag := false
	for i, v := range item {
		if v == 0 {
			if !flag {
				item[0] = i + 1
				flag = true
			} else {
				item[1] = i + 1
			}
		}
	}
	return item[0:2]
}
