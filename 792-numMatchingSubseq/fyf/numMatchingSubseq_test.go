package fyf

import (
	"testing"
)

func TestNumMatchingSubseq(t *testing.T) {
	ret := numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"})
	t.Log(ret)
}
