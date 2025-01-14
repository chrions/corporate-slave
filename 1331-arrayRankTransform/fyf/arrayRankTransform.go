package fyf

import "sort"

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	ret := make([]int, len(arr))
	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	m := map[int]int{}
	sort.Ints(arr1)
	index := 1
	for i := 0; i < len(arr1); i++ {
		if _, ok := m[arr1[i]]; ok {
			continue
		}
		m[arr1[i]] = index
		index++
	}
	for i := 0; i < len(arr); i++ {
		if v, ok := m[arr[i]]; ok {
			ret[i] = v
		}
	}

	return ret
}
