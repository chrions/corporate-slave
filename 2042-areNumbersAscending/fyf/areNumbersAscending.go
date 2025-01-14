package fyf

import (
	"math"
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	current := math.MinInt
	ss := strings.Split(s, " ")
	for i := 0; i < len(ss); i++ {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			continue
		}
		if n > current {
			current = n
		} else {
			return false
		}

	}
	return true
}
