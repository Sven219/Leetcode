package main

import "fmt"

func twoSum(nums []int, target int) []int {
	prevNums := map[int]int{}
	for i, num := range nums {
		targetNum := target - num
		targetNumIndex, ok := prevNums[targetNum]
		if ok {
			return []int{targetNumIndex, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 8
	fmt.Println(twoSum(nums, target))
}
