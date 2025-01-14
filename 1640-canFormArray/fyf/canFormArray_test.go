package fyf

import "testing"

func TestCanFormArray(t *testing.T) {
	ret := canFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}})
	t.Log(ret)
}
