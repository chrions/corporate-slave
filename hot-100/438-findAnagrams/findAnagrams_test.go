package fyf

import (
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	ret := findAnagrams("abbbab", "ab")
	t.Log(ret)
}
