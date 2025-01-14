package fyf

import "strings"

func maxRepeating(sequence string, word string) int {
	isContain := strings.Contains(sequence, word)
	if !isContain {
		return 0
	}
	k := 1
	word1 := word
	for {
		word1 = word1 + word
		isContain := strings.Contains(sequence, word1)
		if !isContain {
			return k
		}
		k++
	}
}
