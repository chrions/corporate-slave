package fyf

func minOperations(boxes string) []int {
	ret := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		step := 0
		for j := i + 1; j < len(boxes); j++ {
			if boxes[j] == '1' {
				step += j - i
			}
		}
		for j := i - 1; j >= 0; j-- {
			if boxes[j] == '1' {
				step += i - j
			}
		}
		ret[i] = step
	}
	return ret
}
