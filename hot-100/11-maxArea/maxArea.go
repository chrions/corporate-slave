package fyf

// 若向内 移动短板 ，水槽的短板 min(h[i],h[j]) 可能变大，因此下个水槽的面积 可能增大 。
// 若向内 移动长板 ，水槽的短板 min(h[i],h[j]) 不变或变小，因此下个水槽的面积 一定变小 。
// 因此，初始化双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新面积最大值，直到两指针相遇时跳出；即可获得最大面积。
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for right > left {
		x := right - left
		y := min(height[right], height[left])
		current := x * y
		if max < current {
			max = current
		}
		if height[right] >= height[left] {
			left++
		} else {
			right--
		}
	}
	return max
}

func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}
