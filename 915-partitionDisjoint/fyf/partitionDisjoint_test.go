package fyf

import "testing"

func TestPartitionDisjoint(t *testing.T) {
	ret := partitionDisjoint([]int{30, 57, 9, 79, 49, 67, 11, 4, 42, 43, 7, 21, 78, 70, 46, 91, 94, 89, 95, 82})
	t.Log(ret)
}
