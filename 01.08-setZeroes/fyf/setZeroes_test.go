package fyf

import "testing"

func TestSetZeroes(t *testing.T) {
	n := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(n)
}
