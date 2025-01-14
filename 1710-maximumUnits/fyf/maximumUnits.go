package fyf

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	num := 0
	truck := 0
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, v := range boxTypes {
		truck += v[0]
		if truck > truckSize {
			num += (truckSize - (truck - v[0])) * v[1]
			break
		}
		num += v[0] * v[1]
	}
	return num
}
