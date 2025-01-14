package fyf

import (
	"testing"
)

func TestReformat(t *testing.T) {
	ret := reformat("abc123")
	t.Log(ret)
}
