package leetcode

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)

	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		for l, r := i + 1, len(nums) - 1; l < r; {
			if l > i + 1 && nums[l] == nums[l - 1] {
				l++
                continue
			}
			if r < len(nums) - 1 && nums[r] == nums[r + 1] {
				r--
                continue
			}

			if nums[i] + nums[l] + nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if nums[i] + nums[l] + nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}