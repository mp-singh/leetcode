package main

import (
	"log"
	_ "strings"
)

func NumJewelsInStones(J string, S string) int {
	count := 0
	for _, i := range J {
		for _, j := range S {
			if i == j {
				count++
			}
		}
	}
	return count
}

func main() {
	i := NumJewelsInStones("aA", "aAAbbbb")
	log.Println(i)
}