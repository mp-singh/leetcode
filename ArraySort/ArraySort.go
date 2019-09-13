package main

import "log"

func SortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	pivotIndex := 1

	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	SortArray(nums[:left])
	SortArray(nums[left+1:])

	return nums
}

func main() {
	sortedArray := SortArray([]int{5, 1, 1, 2, 0, 0})
	log.Println(sortedArray)
}
