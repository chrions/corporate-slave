package fyf

func reorderSpaces(text string) string {
	letters := make([]string, 0)
	space := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			space++
		} else {
			for j := i; i < len(text); j++ {
				if text[j] == ' ' {
					letters = append(letters, text[i:j])
					i = j - 1
					break
				}
				if j == len(text)-1 {
					letters = append(letters, text[i:])
					goto help
				}
			}
		}
	}
help:
	letter := len(letters)
	tail := ""
	if letter == 1 {
		if space == 0 {
			return text
		} else {
			for i := 0; i < space; i++ {
				tail += " "
			}
			return letters[0] + tail
		}

	}

	n := space / (letter - 1)
	spaces := ""
	for i := 0; i < n; i++ {
		spaces += " "
	}
	m := space % (letter - 1)
	for i := 0; i < m; i++ {
		tail += " "
	}
	s := ""
	for i, v := range letters {
		if i == len(letters)-1 {
			s += v
			break
		}
		s += v + spaces
	}
	return s + tail
}
