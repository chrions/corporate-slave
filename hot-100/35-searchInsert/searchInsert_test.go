package fyf

import "testing"

func TestSearchInsert(t *testing.T) {
	array := []int{1, 3, 5, 6}
	ret := searchInsert(array, 2)
	t.Log(ret)
}
