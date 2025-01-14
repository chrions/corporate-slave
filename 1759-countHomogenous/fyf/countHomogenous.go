package fyf

const a = 1e9 + 7

func countHomogenous(s string) int {
	equal := ""
	sub := 0
	ret := 0

	for i := 0; i < len(s); i++ {
		if len(equal) <= 1 {
			equal = string(s[i])
		}
		if i+1 < len(s) && s[i+1] == s[i] {
			equal += string(s[i+1])
		}
		if len(equal) > 1 && i+1 < len(s) && s[i+1] != s[i] {
			sub += len(equal)
			ret += (len(equal) + 1) * len(equal) / 2
			equal = ""
		}
		if len(equal) > 1 && i == len(s)-1 {
			sub += len(equal)
			ret += (len(equal) + 1) * len(equal) / 2
		}
	}

	ret = ret + len(s) - sub

	return ret % a
}
