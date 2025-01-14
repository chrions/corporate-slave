package fyf

import (
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	sort.Slice(items1, func(i, j int) bool {
		return items1[i][0] < items1[j][0]
	})
	sort.Slice(items2, func(i, j int) bool {
		return items2[i][0] < items2[j][0]
	})
	ret := make([][]int, 0)
	i, j := 0, 0
	for i < len(items1) && j < len(items2) {
		if items1[i][0] == items2[j][0] {
			temp := []int{items1[i][0], items1[i][1] + items2[j][1]}
			ret = append(ret, temp)
			i++
			j++
			continue
		}
		if items1[i][0] > items2[j][0] {
			ret = append(ret, items2[j])
			j++
			continue
		}
		if items1[i][0] < items2[j][0] {
			ret = append(ret, items1[i])
			i++
			continue
		}
	}

	if j == len(items2) {
		ret = append(ret, items1[i:]...)
	}

	if i == len(items1) {
		ret = append(ret, items2[j:]...)
	}
	return ret
}
