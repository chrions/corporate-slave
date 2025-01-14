package main

func scoreOfParentheses(s string) int {
	stack := []int{0}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else {
			v := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			stack[len(stack)-1] += max(v*2, 1)
		}
	}
	return stack[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
