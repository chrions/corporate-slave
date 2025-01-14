package fyf

func canFormArray(arr []int, pieces [][]int) bool {
	for _, v := range pieces {
		ret := ContainsArr(arr, v)
		if ret == false {
			return false
		}
	}
	return true
}

func ContainsArr(src []int, dest []int) bool {
	head := dest[0]
	flag := false
	for i, v := range src {
		if v == head {
			flag = true
			for j := 0; j < len(dest); j++ {
				if i >= len(src) {
					return false
				}
				if dest[j] != src[i] {
					return false
				}
				i++
			}
			break
		}
	}
	return flag
}
