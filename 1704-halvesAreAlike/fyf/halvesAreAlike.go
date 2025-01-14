package fyf

var m = map[uint8]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func halvesAreAlike(s string) bool {
	s1 := s[0 : len(s)/2]
	s2 := s[len(s)/2:]

	n1 := 0
	for _, v := range s1 {
		if _, ok := m[uint8(v)]; ok {
			n1++
		}
	}

	n2 := 0
	for _, v := range s2 {
		if _, ok := m[uint8(v)]; ok {
			n2++
		}
	}
	if n1 == n2 {
		return true
	}
	return false
}
