package fyf

import "testing"

func TestShuffle(t *testing.T) {
	ret := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	t.Log(ret)
}
