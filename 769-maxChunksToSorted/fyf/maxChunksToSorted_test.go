package fyf

import "testing"

func TestMaxChunksToSorted(t *testing.T) {
	ret := maxChunksToSorted([]int{2, 1, 3, 4, 4})
	t.Log(ret)
}
