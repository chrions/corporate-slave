package fyf

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	a := 1
	b := 2
	c := 0
	for i := 3; i < n+1; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
