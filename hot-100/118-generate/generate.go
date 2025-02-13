package fyf

func generate(numRows int) [][]int {
	ret := make([][]int, 0, numRows)

	ret = append(ret, []int{1})
	for i := 1; i < numRows; i++ {
		last := ret[len(ret)-1]
		temp := make([]int, len(last)+1)
		temp[0], temp[len(temp)-1] = 1, 1
		for j := 1; j < len(temp)-1; j++ {
			temp[j] = last[j-1] + last[j]
		}
		ret = append(ret, temp)
	}
	return ret
}
