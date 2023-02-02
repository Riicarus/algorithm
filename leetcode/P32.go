package leetcode

import "fmt"

func LongestValidParentheses(s string) int {
	count := 0

	for idx := 0; idx < len(s); idx++ {
		fmt.Println(idx)
		if s[idx] != '(' {
			continue
		}

		l := 0
		for j := idx; j < len(s); j++ {
			// (, l自增
			if s[j] == '(' {
				l++
			} else {
				// ), 需要有可以匹配的 (
				if l > 0 {
					l--
					if l == 0 {
						if j-idx+1 > count {
							count = j - idx + 1
						}
					}
				} else {
					// 遇到 ), 且可以匹配的 ( 个数为 0, ) 盈余, 从下一个位置继续匹配
					idx = j + 1
				}
			}
		}
	}

	return count
}
