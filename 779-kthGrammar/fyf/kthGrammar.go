package fyf

import (
	"strconv"
)

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	s := "0"
	for i := 2; i <= n; i++ {
		s = s + help(s)
	}
	m := s[k-1]
	mm, _ := strconv.Atoi(string(m))
	return mm
}

func help(s string) string {
	ss := make([]uint8, len(s))
	for i, v := range s {
		if v == '0' {
			ss[i] = '1'
		}
		if v == '1' {
			ss[i] = '0'
		}
	}
	return string(ss)
}

func kthGrammar1(n int, k int) int {
	if n == 1 {
		return 0
	}
	if kthGrammar(n-1, (k+1)/2) == 0 {
		return (k%2 + 1) % 2
	}
	return k % 2
}
