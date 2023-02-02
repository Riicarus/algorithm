package leetcode

func MaxArea(height []int) int {
	res := 0

	l, r := 0, len(height) - 1

	for ; l < r; {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}