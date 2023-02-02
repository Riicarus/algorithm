package leetcode

import (
	"fmt"
	"sort"
)

func NextPermutation(nums []int)  {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			for j := len(nums) - 1; j > i - 1; j-- {
				if nums[j] > nums[i - 1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					sort.Ints(nums[i:])
					fmt.Println(nums)
					return
				}
			}
		}
	}
	
	sort.Ints(nums)
	fmt.Println(nums)
}
