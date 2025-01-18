package fyf

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	// 如果第一项就能直接跳到最后一项
	if nums[0]+0 >= len(nums)-1 {
		return 1
	}

	start := 0
	end := nums[0]
	step := 1 // 第一次跳跃已确定

	// 开始从第二步跳跃
	for start <= end {
		// 找到当前区间能跳到的最远位置
		position := maxStep(nums, start, end)

		// 如果当前位置+能跳的最大步数可以到达终点，跳跃一次就能到达终点
		if position+nums[position] >= len(nums)-1 {
			step++
			break
		}

		// 更新起始位置和结束位置，准备下一步跳跃
		start = position
		end = position + nums[position]
		step++
	}

	return step
}

// maxStep 找到当前跳跃范围内，能够跳跃到的最远位置
func maxStep(nums []int, start, end int) int {
	maxReach := -1
	position := start
	for i := start; i <= end; i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
			position = i
		}
	}
	return position
}
