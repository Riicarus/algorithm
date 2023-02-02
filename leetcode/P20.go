package leetcode

func IsValid(s string) bool {
	if len(s)%2 == 1 || len(s) == 0 {
		return false
	}

	stack := make([]rune, 10000)

	top := -1

	for _, c := range s {
		switch c {
		case '(':
			top++
			stack[top] = c
		case '[':
			top++
			stack[top] = c
		case '{':
			top++
			stack[top] = c
		case ')':
			if stack[top] != '(' {
				return false
			}
			top--
		case ']':
			if stack[top] != '[' {
				return false
			}
			top--
		case '}':
			if stack[top] != '{' {
				return false
			}
			top--
		}
	}

	return top == 0
}
