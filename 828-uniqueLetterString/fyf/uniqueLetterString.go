package fyf

func uniqueLetterString(s string) (ans int) {
	for i, v := range s {
		l, r := 0, 0
		for j := i - 1; j >= 0; j-- {
			if rune(s[j]) != v {
				l++
			} else {
				break
			}
		}
		for j := i + 1; j < len(s); j++ {
			if rune(s[j]) != v {
				r++
			} else {
				break
			}
		}
		ans += 1 + l + r + (l * r)
	}
	return
}
