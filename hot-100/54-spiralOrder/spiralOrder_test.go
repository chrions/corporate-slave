package fyf

import "testing"

func TestSpiralOrder(t *testing.T) {
	ret := spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	})
	t.Log(ret)
}
