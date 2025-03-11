package fyf

func isValid(s string) bool {
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := []uint8{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if v, ok := m[string(s[i])]; ok {
			top := stack[len(stack)-1]
			if string(top) == v {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
