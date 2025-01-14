package fyf

func maxEqualFreq(nums []int) int {
	n, ret := len(nums), 0
	freq, total := make(map[int]int), make(map[int]int)

	ma, size := 0, 0
	for i := 0; i < n; i++ {
		if freq[nums[i]] == 0 {
			size++
		}
		if freq[nums[i]] >= 1 {
			total[freq[nums[i]]]--
		}
		freq[nums[i]]++
		ma = maxInt(freq[nums[i]], ma)
		total[freq[nums[i]]]++

		//=== 1.多种数字，每个数字各出现1次
		//=== 2.只有1个数字
		//=== 3.多种数字，频率为ma的有1个，其余频率均为ma-1
		//=== 4.多种数字，频率为ma的有n-1个，频率为1的有1个
		if total[1] == i+1 || freq[nums[i]] == i+1 || (total[ma] == 1 && total[ma-1] == size-1) || (total[1] == 1 && total[ma] == size-1) {
			ret = maxInt(ret, i+1)
		}
	}
	return ret
}
func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
