package main

import "log"

func sortArrayByParity(a []int) []int {
	result := make([]int, len(a))
	i, j := 0, len(a)-1
	for _, e := range a {
		if e%2 == 0 {
			result[i] = e
			i++
			continue
		}
		result[j] = e
		j--
	}
	return result
}

func main() {
	result := sortArrayByParity([]int{3, 1, 2, 4})
	log.Println(result)
}
