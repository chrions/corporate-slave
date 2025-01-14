package fyf

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	ret := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ret[len(ret)-1][1] {
			ret[len(ret)-1][1] = max(ret[len(ret)-1][1], intervals[i][1])
			continue
		}
		if intervals[i][0] > ret[len(ret)-1][1] {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
