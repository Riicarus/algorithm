package leetcode

func Trap(height []int) int {
	n := len(height)

	if n == 0 {
		return 0
	}

	lmax := make([]int, n)
	rmax := make([]int, n)

	lmax[0] = height[0]
	rmax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}

	f := make([]int, n)
	for i, h := range height {
		if i == 0 {
			f[i] = 0
		} else {
			f[i] = f[i-1] + min(lmax[i], rmax[i]) - h
		}
	}

	return f[n-1]
}