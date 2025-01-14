package fyf

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	temp := make([]int, 0, len(m))
	for k, _ := range m {
		temp = append(temp, k)
	}
	sort.Ints(temp)
	max := 0
	current := 1
	for i := 0; i < len(temp); i++ {
		if i+1 < len(temp) && temp[i+1] == temp[i]+1 {
			current++
			continue
		}
		if current > max {
			max = current
		}
		current = 1
	}
	return max
}
