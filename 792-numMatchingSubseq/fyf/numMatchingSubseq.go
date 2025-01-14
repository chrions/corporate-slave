package fyf

import "strings"

func numMatchingSubseq(s string, words []string) int {
	ret := len(words)
	for _, v := range words {
		temp := s
		for i := 0; i < len(v); i++ {
			r := strings.IndexRune(temp, rune(v[i]))
			if r == -1 {
				ret--
				break
			}
			temp = temp[r+1:]
		}

	}
	return ret
}
