package fyf

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	generate(n, n, "", &ret)
	return ret
}

func generate(left, right int, s string, ret *[]string) {
	if right == 0 {
		*ret = append(*ret, s)
	}
	if left > 0 {
		generate(left-1, right, s+"(", ret)
	}
	if right > left {
		generate(left, right-1, s+")", ret)
	}
}
