package fyf

import "testing"

func TestArray(t *testing.T) {
	arr := []int{100, 100, 100}
	ret := arrayRankTransform(arr)
	t.Log(ret)
}
