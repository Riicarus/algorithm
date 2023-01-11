package leetcode

import "sort"

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func MinimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])

	var dfs func(r, c, limit int, seen [][]bool) bool

	dfs = func(curR, curC, limit int, seen [][]bool) bool {
		// 终止条件: 到达终点
		if curR == m-1 && curC == n-1 {
			return true
		}

		// 将该点设置为已经访问
		seen[curR][curC] = true
		for _, d := range dirs {
			// 向上下左右访问
			r, c := curR + d[0], curC + d[1]
			if r >= 0 && r < m && c >= 0 && c < n && !seen[r][c] &&
				abs(heights[r][c] - heights[curR][curC]) <= limit {
				if dfs(r, c, limit, seen) {
					return true
				}
			}
		}
		// 这里不能把 seen[curR][curC] 重置为 false：
		// 如果(curR, curC) 点能在 limit 限制下到达终点会返回 true，不能到达会返回 false，后边也不用再来尝试这个点
		return false
	}


	return sort.Search(1e6, func(i int) bool {
		return dfs(0, 0, i, genMemo(m, n))
	})
}

func genMemo(m, n int) [][]bool {
	res := make([][]bool, m)
	for i := range res {
		res[i] = make([]bool, n)
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}