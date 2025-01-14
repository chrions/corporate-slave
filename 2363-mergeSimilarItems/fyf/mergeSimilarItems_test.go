package fyf

import (
	"testing"
)

func TestMergeSimilarItems(t *testing.T) {
	ret := mergeSimilarItems([][]int{{5, 1}, {4, 2}, {3, 3}, {2, 4}, {1, 5}},
		[][]int{{7, 1}, {6, 2}, {5, 3}, {4, 4}})
	t.Log(ret)
}
