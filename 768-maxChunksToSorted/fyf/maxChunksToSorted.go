package fyf

func maxChunksToSorted2(arr []int) int {
	stack := make([]int, 0)
	for _, v := range arr {
		if len(stack) == 0 || v >= stack[len(stack)-1] {
			stack = append(stack, v)
		} else {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for len(stack) > 0 && stack[len(stack)-1] > v {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, n)
		}
	}
	return len(stack)
}
