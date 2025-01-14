package fyf

import "fmt"

func getFolderNames(names []string) []string {
	ret := make([]string, len(names))
	suffix := map[string]int{}
	for i, name := range names {
		n := suffix[name]
		if n == 0 {
			suffix[name] = 1
			ret[i] = name
			continue
		}
		for suffix[fmt.Sprintf("%s(%d)", name, n)] > 0 {
			n++
		}
		ret[i] = fmt.Sprintf("%s(%d)", name, n)
		suffix[ret[i]] = 1
		suffix[name] = n + 1
	}
	return ret
}
