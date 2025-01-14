package fyf

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxLength := 0
	j := 0
	for i := 0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok {
			if index < j { //xxcxxaxxaxc j前面的重复字符不能算了
				m[s[i]] = i
				continue
			}
			maxLength = max(maxLength, i-j)
			delete(m, s[i])
			j = index + 1
		}
		m[s[i]] = i
	}
	maxLength = max(maxLength, len(s)-j) //abcdefg 没有重复的情况
	return maxLength
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
