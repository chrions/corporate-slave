package fyf

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	ret := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	t.Log(ret)
}
