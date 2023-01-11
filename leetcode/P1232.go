package leetcode

func CheckStraightLine(coordinates [][]int) bool {
	for i := 1; i < len(coordinates) - 1; i++ {

		if !((coordinates[i][1] - coordinates[0][1]) * (coordinates[i + 1][0] - coordinates[0][0]) == (coordinates[i + 1][1] - coordinates[0][1]) * (coordinates[i][0] - coordinates[0][0])) {
			return false
		}
	}

	return true
}