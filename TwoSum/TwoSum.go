package main

import "log"

func TwoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	indexes := TwoSum([]int{3, 2, 4}, 6)
	log.Println(indexes)
}
