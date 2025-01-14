package fyf

func maxChunksToSorted(arr []int) int {
	stack := []int{}
	for _, v := range arr {
		if len(stack) == 0 || v >= stack[len(stack)-1] {
			stack = append(stack, v)
		} else {
			mx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for len(stack) > 0 && stack[len(stack)-1] > v {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, mx)
		}
	}
	return len(stack)
}
