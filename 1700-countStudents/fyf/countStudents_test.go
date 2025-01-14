package fyf

import "testing"

func TestCountStudents(t *testing.T) {
	ret := countStudents([]int{0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1}, []int{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0})
	t.Log(ret)
}
