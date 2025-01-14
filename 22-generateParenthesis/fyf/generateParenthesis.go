package fyf

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	generate(n, n, "", &ret)
	return ret
}

func generate(left, right int, s string, ret *[]string) {
	/*
	   回溯跳出条件，
	   并不需要判断左括号是否用完，因为右括号生成的条件 right > left ，
	   所以右括号用完了就意味着左括号必定用完了
	*/
	if right == 0 {
		*ret = append(*ret, s)
	}

	// 生成左括号
	if left > 0 {
		generate(left-1, right, s+"(", ret)
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		generate(left, right-1, s+")", ret)
	}
}
