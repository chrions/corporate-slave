package fyf

import (
	"testing"
)

func TestSetZeroes(t *testing.T) {
	setZeroes([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	})
}
