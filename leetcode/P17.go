package leetcode

var digitLetterMap = make(map[byte][]string)

func InitMap() {
	digitLetterMap['1'] = nil
	digitLetterMap['2'] = []string{"a", "b", "c"}
	digitLetterMap['3'] = []string{"d", "e", "f"}
	digitLetterMap['4'] = []string{"g", "h", "i"}
	digitLetterMap['5'] = []string{"j", "k", "l"}
	digitLetterMap['6'] = []string{"m", "n", "o"}
	digitLetterMap['7'] = []string{"p", "q", "r", "s"}
	digitLetterMap['8'] = []string{"t", "u", "v"}
	digitLetterMap['9'] = []string{"w", "x", "y", "z"}
}

var results = []string{}

func LetterCombinations(digits string) []string {
	InitMap()

	results = make([]string, 0)

	if len(digits) == 0 {
		return []string{}
	}

	backTrack(digits, 0, "")

	return results
}

func backTrack(digits string, index int, res string) {
	if index == len(digits) {
		results = append(results, res)
		return
	}

	letters := digitLetterMap[digits[index]]

	for _, l := range letters {
		backTrack(digits, index + 1, res + l)
	}
}