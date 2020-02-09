package main

import (
	"log"
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i+=2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

func main() {
	indexes := singleNumber([]int{4,3, 3, 2, 2})
	log.Println(indexes)
}
