package fyf

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	pre := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		cur := make([]int, i+1)
		cur[0] = 1
		cur[i] = 1
		for j := 0; j < len(pre); j++ {
			if j+1 < len(pre) {
				cur[j+1] = pre[j] + pre[j+1]
			}
		}
		pre = cur
	}
	return pre
}
