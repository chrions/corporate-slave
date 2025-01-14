package fyf

func arraySign(nums []int) int {
	ret := 1
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v > 0 {
			ret = ret
		}
		if v < 0 {
			ret = -ret
		}
	}
	return ret
}
