package fyf

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, v := range operations {
		switch v {
		case "++X", "X++":
			x++
		case "--X", "X--":
			x--
		}
	}
	return x
}
