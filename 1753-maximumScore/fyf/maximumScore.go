package fyf

import (
	"sort"
)

func maximumScore(a int, b int, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	ret := 0
	if arr[0]+arr[1] <= arr[2] {
		ret = arr[0] + arr[1]
	} else {
		for {
			arr[0], arr[1] = arr[0]-1, arr[1]-1
			ret++
			if arr[0]+arr[1] <= arr[2] {
				ret += arr[0] + arr[1]
				break
			}
		}

	}

	return ret
}
