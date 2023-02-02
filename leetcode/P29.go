package leetcode

import (
	"fmt"
)

func Devide(dividend, divisor int) int {
	// out of range handler
	if (divisor == 1 && dividend > (2<<30 - 1)) || ((divisor == -1) && (dividend < -(2<<30 - 1))) || (divisor == 1 && dividend < -(2<<30)) || (divisor == -1 && dividend > (2<<30)){
		return 2<<30 - 1
	}

	// set all number to positive
	negative := false
	if dividend < 0 {
		dividend = -dividend
		negative = !negative
	}
	if divisor < 0 {
		divisor = -divisor
		negative = !negative
	}

	if negative {
		return -div(dividend, divisor)
	}

	return div(dividend, divisor)
}

func div(dividend, divisor int) int {
	if dividend < divisor {
		return 0
	}

	stageRes := 1
	nextDivisor := divisor
	for ;(nextDivisor + nextDivisor) <= dividend; {
		stageRes = stageRes + stageRes
		nextDivisor = nextDivisor + nextDivisor
	}

	return stageRes + div(dividend - nextDivisor, divisor)
}

func Main() {
	dividend := 0
	fmt.Scan(&dividend)

	divisor := 0
	fmt.Scan(&divisor)

	fmt.Println(Devide(dividend, divisor))
}