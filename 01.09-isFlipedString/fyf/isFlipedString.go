package fyf

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s1+s1, s2)
}
