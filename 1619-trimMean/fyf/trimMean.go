package fyf

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	length := len(arr)
	n := int(float64(length) * 0.05)
	newArr := arr[n : length-n]

	total := 0
	for _, v := range newArr {
		total += v
	}
	return float64(total) / float64(len(newArr))
}
