package fyf

func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		switch presses {
		case 1:
			return 3
		default:
			return 4
		}
	} else {
		switch presses {
		case 1:
			return 4
		case 2:
			return 7
		default:
			return 8
		}
	}
}
