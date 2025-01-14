package fyf

import "testing"

func TestFind(t *testing.T) {
	ret := findClosestElements([]int{1, 3}, 1, 2)
	t.Log(ret)
}
