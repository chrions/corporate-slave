package fyf

func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	left := 0
	ans := 0
	for right, x := range fruits {
		cnt[x]++
		for len(cnt) > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
