package leetcode

func GenerateParenthesis(n int) []string {
	res := make([]string, 0)

	inCount := n
	outCount := 0

	resString := ""

	nextStep(inCount, outCount, resString, &res)

	return res
}

func nextStep(inCount int, outCount int, resString string, res *[]string) {
	if inCount == 0 && outCount == 0 {
		*res = append(*res, resString)
		return
	}
	
	// (
	if inCount > 0 {
		nextStep(inCount - 1, outCount + 1, resString + "(", res)
	}

	// )
	if outCount > 0 {
		nextStep(inCount, outCount - 1, resString + ")", res)
	}
}
