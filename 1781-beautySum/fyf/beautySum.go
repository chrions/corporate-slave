package fyf

import "math"

func beautySum(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		cnt := [26]int{}
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a']++
			min, max := math.MaxInt32, math.MinInt32
			for _, v := range cnt {
				if v > 0 && v > max {
					max = v
				}
				if v > 0 && v < min {
					min = v
				}
			}
			ret += max - min
		}
	}
	return ret
}
