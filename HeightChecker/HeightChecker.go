package main

import (
	"log"
	"sort"
)

func heightChecker(heights []int) int {
	temp := make([]int, len(heights))
	copy(temp, heights)
	sort.Ints(temp)
	count := 0
	for i := range heights {
		if heights[i] != temp[i] {
			count++
		}
	}
	return count
}

func main() {
	result := heightChecker([]int{1,1,4,2,1,3})
	log.Println(result)
}
