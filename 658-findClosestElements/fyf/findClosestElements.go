package fyf

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[0:k]
	}
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	ret := getNumbers(arr, k, x)
	sort.Ints(ret)
	return ret
}

func getNumbers(arr []int, k int, x int) []int {
	index := 0
	ret := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) && arr[i] == x && arr[i+1] > x {
			index = i
			break
		}
		if i+1 < len(arr) && arr[i] < x && arr[i+1] > x {
			if abs(arr[i]-x) <= abs(arr[i+1]-x) {
				index = i
				break
			} else {
				index = i + 1
				break
			}
		}
	}
	ret = append(ret, arr[index])
	i := index - 1
	j := index + 1
	for i >= 0 && j < len(arr) && len(ret) < k {
		if abs(arr[i]-x) == abs(arr[j]-x) {
			ret = append(ret, arr[i])
			i = i - 1
			continue
		}
		if abs(arr[i]-x) < abs(arr[j]-x) {
			ret = append(ret, arr[i])
			i = i - 1
		} else {
			ret = append(ret, arr[j])
			j = j + 1
		}

	}
	if i < 0 && len(ret) < k {
		sort.Ints(ret)
		n := ret[len(ret)-1]
		maxI := 0
		for i := 0; i < len(arr); i++ {
			if i+1 < len(arr) && arr[i] == n && arr[i+1] > n {
				maxI = i
				break
			}
		}
		ret = append(ret, arr[maxI+1:maxI+1+(k-len(ret))]...)
	}
	if j > len(arr)-1 && len(ret) < k {
		sort.Ints(ret)
		n := ret[0]
		minI := 0
		for i := len(arr) - 1; i >= 0; i-- {
			if i-1 >= 0 && arr[i] == n && arr[i-1] < n {
				minI = i
				break
			}
		}

		ret = append(ret, arr[minI-(k-len(ret)):minI]...)
	}
	return ret
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
