package fyf

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := ""
	for _, v := range word1 {
		s1 += v
	}
	s2 := ""
	for _, v := range word2 {
		s2 += v
	}
	if s1 == s2 {
		return true
	}
	return false
}
