package fyf

import (
	"testing"
)

func TestReorderSpaces(t *testing.T) {
	ret := reorderSpaces("  a")
	t.Log(ret)
}
