package fyf

import (
	"testing"
)

func TestJump(t *testing.T) {
	nums := []int{6, 9, 1, 5, 6, 0, 0, 5, 9}
	ret := jump(nums)
	t.Log(ret)
}
