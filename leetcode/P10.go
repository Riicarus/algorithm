package leetcode

// 从右往左匹配
func IsMatch2(s string, p string) bool {
	idx := len(s) - 1
	for i := len(p) - 1; i >= 0 && idx >= 0; i-- {

		// 不是通配符的判断
		if p[i] != '.' && p[i] != '*' {
			// 判断 b*b
			if i < len(p) - 2 && p[i + 1] == '*' && p[i + 2] == p[i] {
				continue
			}

			if p[i] != s[idx] {
				return false
			}
			idx--
			continue
		}

		// 匹配 .
		if p[i] == '.' {
			idx--
			continue
		}

		// 匹配 *
		if p[i] == '*' {
			pre := p[i-1]
			// 特判 .*, 匹配所有
			if pre == '.' {
				if i == 0 {
					return true
				}
				continue
			}

			for j := 0; idx - j >= 0 && s[idx - j] == pre; j++ {
				idx--
			}
			idx--
		}
	}

	return idx < 0
}

func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
