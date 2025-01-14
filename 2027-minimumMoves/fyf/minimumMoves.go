package fyf

func minimumMoves(s string) int {
	n := -1
	ret := 0
	for i, v := range s {
		if v == 'X' && i > n {
			ret++
			n = i + 2
		}
	}
	return ret
}
