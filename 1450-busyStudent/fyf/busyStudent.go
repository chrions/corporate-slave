package fyf

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ret := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ret++
		}
	}
	return ret
}
