package fyf

import "testing"

func TestTotalFruit(t *testing.T) {
	ret := totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4})
	t.Log(ret)
}
