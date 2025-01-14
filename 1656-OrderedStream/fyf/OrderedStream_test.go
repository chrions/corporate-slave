package fyf

import "testing"

func TestOrderedStream_Insert(t *testing.T) {
	s := Constructor(5)
	ret := s.Insert(3, "ccccc")
	t.Log(ret)
	ret = s.Insert(1, "aaaaa")
	t.Log(ret)
	ret = s.Insert(2, "bbbbb")
	t.Log(ret)
	ret = s.Insert(5, "eeeee")
	t.Log(ret)
	ret = s.Insert(4, "ddddd")
	t.Log(ret)
}
