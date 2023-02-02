package main

import (
	"fmt"

	_ "github.com/riicarus/algorithm/acwing"
	"github.com/riicarus/algorithm/leetcode"
)

func main() {
	// acwing.Search([]int{1, 2, 2, 3, 3, 4}, []int{3, 4})

	// fmt.Printf("leetcode.LongestPalindrome(\"babad\"): %v\n", leetcode.LongestPalindrome("babad"))
	// fmt.Printf("leetcode.LongestPalindrome(\"babab\"): %v\n", leetcode.LongestPalindrome("babab"))
	// fmt.Printf("leetcode.LongestPalindrome(\"baba\"): %v\n", leetcode.LongestPalindrome("baba"))

	// fmt.Printf("leetcode.IsMatch2(\"aa\", \"a\"): %v\n", leetcode.IsMatch2("aa", "a"))
	// fmt.Printf("leetcode.IsMatch2(\"aa\", \"a*\"): %v\n", leetcode.IsMatch2("aa", "a*"))
	// fmt.Printf("leetcode.IsMatch2(\"aab\", \"c*a*b\"): %v\n", leetcode.IsMatch2("aab", "c*a*b"))
	// fmt.Printf("leetcode.IsMatch2(\"aabbbc\", \"a.b*bc\"): %v\n", leetcode.IsMatch2("aabbbc", "a.b*bc"))
	
	// fmt.Printf("leetcode.LetterCombinations(\"\"): %v\n", leetcode.LetterCombinations(""))
	// fmt.Printf("leetcode.LetterCombinations(\"23\"): %v\n", leetcode.LetterCombinations("23"))

	// fmt.Printf("leetcode.GenerateParenthesis(3): %v\n", leetcode.GenerateParenthesis(3))

	// fmt.Printf("leetcode.LongestValidParentheses(\"(()\"): %v\n", leetcode.LongestValidParentheses("(()"))
	// fmt.Printf("leetcode.LongestValidParentheses(\"()()\"): %v\n", leetcode.LongestValidParentheses("()()"))
	// fmt.Printf("leetcode.LongestValidParentheses(\"()(()\"): %v\n", leetcode.LongestValidParentheses("()(()"))
	// fmt.Printf("leetcode.LongestValidParentheses(\")()())()()(\"): %v\n", leetcode.LongestValidParentheses(")()())()()("))
	// fmt.Printf("leetcode.LongestValidParentheses(\"(())(\"): %v\n", leetcode.LongestValidParentheses("(())("))

	fmt.Printf("leetcode.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}): %v\n", leetcode.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}