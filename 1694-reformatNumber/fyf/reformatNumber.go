package fyf

import "strings"

func reformatNumber(number string) string {
	s := strings.Replace(strings.Replace(number, "-", "", -1), " ", "", -1)
	n := len(s)
	if n < 4 {
		return s
	}
	ss := ""
	for n > 4 {
		ss += s[0:3] + "-"
		s = s[3:]
		n = len(s)
	}

	if n == 4 {
		ss += s[0:2] + "-" + s[2:]
		return ss
	}

	ss += s
	return ss
}
