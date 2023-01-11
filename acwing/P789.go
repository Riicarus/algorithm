package acwing

import "fmt"

// type: binary search
func Search(sarr []int, qarr []int) {
	s, e := 0, 0

	for _, q := range qarr {
		l, r := 0, len(sarr) - 1
		for ;l < r; {
			mid := (l + r + 1) >> 1

			if searchLeft(q, sarr[mid]) {
				l = mid
			} else {
				r = mid - 1
			}
		}
		s = l + 1

		l, r = 0, len(sarr) - 1
		for ; l < r; {
			mid := (l + r + 1) >> 1

			if searchRight(q, sarr[mid]) {
				r = mid - 1
			} else {
				l = mid
			}
		}
		e = l

		fmt.Printf("%d %d\n", s, e)
	}
}

func searchLeft(dst, mid int) bool {
	return mid < dst
}

func searchRight(dst, mid int) bool {
	return mid > dst
}