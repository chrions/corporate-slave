package fyf

import "sort"

func reformat(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	letterIndex := 0
	for i := 0; i < len(b); i++ {
		if b[i] >= 97 && b[i] <= 122 {
			letterIndex = i
			break
		}
	}

	letterLen := len(b) - letterIndex
	numLen := letterIndex
	if numLen-letterLen != 0 && numLen-letterLen != -1 && numLen-letterLen != 1 {
		return ""
	}

	ret := make([]byte, 0, len(b))
	flag := true
	j := letterIndex
	k := 0
	for i := 0; i < len(b); i++ {
		if numLen-letterLen > 0 {
			if flag {
				ret = append(ret, b[k])
				k++
				flag = !flag
			} else {
				if j == len(s) {
					ret = append(ret, b[letterIndex-1])
					break
				}
				ret = append(ret, b[j])
				j++
				flag = !flag
			}
		} else {
			if flag {
				ret = append(ret, b[j])
				j++
				flag = !flag
			} else {
				if k == letterLen {
					ret = append(ret, b[len(s)-1])
					break
				}
				ret = append(ret, b[k])
				flag = !flag
				k++
			}

		}
	}
	return string(ret)
}
