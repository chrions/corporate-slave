package fyf

func generateTheString(n int) string {
	if n == 0 {
		return ""
	}

	ret := make([]string, n)
	rest := n % 2
	if rest == 0 {
		for i := 0; i < n; i++ {
			if i == 0 {
				ret[i] = "a"
				continue
			}
			ret = append(ret, "b")
		}
	} else {
		for i := 0; i < n; i++ {
			if i == 0 {
				ret[i] = "c"
				continue
			}
			s := generateTheString(n - 1)
			ret = append(ret, s)
			return arrToStr(ret)
		}
	}
	return arrToStr(ret)
}

func arrToStr(s []string) string {
	ss := ""
	for _, v := range s {
		ss += v
	}
	return ss
}
