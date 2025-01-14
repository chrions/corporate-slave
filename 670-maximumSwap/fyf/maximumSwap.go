package fyf

import (
	"strconv"
)

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	max := s[0]
	maxI := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= max {
			max = s[i]
			maxI = i
		}
	}

	if max != s[0] {
		s[0], s[maxI] = s[maxI], s[0]
		n, _ := strconv.Atoi(string(s))
		return n
	}

	for i := 0; i < len(s); i++ {
		if s[i] < max {
			sec := s[i]
			flag := true
			temp := 0
			temp2 := 0
			for j := i + 1; j < len(s); j++ {
				if s[j] >= sec {
					sec = s[j]
					temp = j
				}
				if flag && s[j] < s[i] && j != len(s)-1 {
					flag = false
					temp2 = j
				}

			}

			if sec > s[i] {
				s[i], s[temp] = s[temp], s[i]
				n, _ := strconv.Atoi(string(s))
				return n
			}
			if sec == s[i] && temp != 0 && temp2 != 0 {
				s[temp2], s[temp] = s[temp], s[temp2]
				n, _ := strconv.Atoi(string(s))
				return n
			}
		}
	}

	return num
}
