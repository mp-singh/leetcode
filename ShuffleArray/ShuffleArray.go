package main

import "log"

func Shuffle(nums []int, n int) []int {
	//result := make([]int, len(nums))
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[n+i])
	}
	return result
}

func main() {
	indexes := Shuffle([]int{2,5,1,3,4,7}, 3)
	log.Println(indexes)
}
