package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if v >= 0 || k == 0 {
			break
		}
		nums[i] = -nums[i]
		k--
	}
	sort.Ints(nums)
	k = k % 2
	ans := 0
	for i, num := range nums {
		if i == 0 && k == 1 {
			num = -num
		}
		ans += num
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 3
	fmt.Println(largestSumAfterKNegations(nums, k))
}
