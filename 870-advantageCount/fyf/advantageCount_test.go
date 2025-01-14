package fyf

import (
	"testing"
)

func TestAdvantageCount(t *testing.T) {
	ret := advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11})
	t.Log(ret)
}
