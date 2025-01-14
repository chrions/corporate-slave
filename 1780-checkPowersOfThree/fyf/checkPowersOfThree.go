package fyf

func checkPowersOfThree(n int) bool {
	for n > 0 {
		m := n % 3
		if m != 1 && m != 0 {
			return false
		}
		n = n / 3
	}
	return true
}
