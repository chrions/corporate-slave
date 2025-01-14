package fyf

import "testing"

func TestMinStartValue(t *testing.T) {
	ret := minStartValue([]int{2, 3, 5, -5, -1})
	t.Log(ret)
}
