package fyf

import (
	"testing"
)

func TestMaxRepeating(t *testing.T) {
	ret := maxRepeating("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba")
	t.Log(ret)
}
