package main

func RemoveDuplicates(nums []int) int {
	j := 0
	for i := range nums {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func main() {
	nums := []int{1, 1, 2}
	size := RemoveDuplicates(nums)
	for i := 0; i < size; i++ {
		print(nums[i])
	}
}
