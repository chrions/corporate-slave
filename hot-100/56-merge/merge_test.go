package fyf

import "testing"

func TestMerge(t *testing.T) {
	ret := merge([][]int{
		{1, 4},
		{0, 2},
		{5, 6},
	})
	t.Log(ret)
}
