package main

import (
	"log"
	"strings"
)

func lengthOfLastWord(s string) int {
	indexOfLastSpace := 0
	s = strings.TrimSpace(s)
	for i, char := range s {
		if char == 32 {
			indexOfLastSpace = i
		}
	}
	if indexOfLastSpace == 0 {
		if s != "" {
			return len(s)
		}
		return 0
	}
	return len(s[indexOfLastSpace+1:])
}

func main() {
	log.Println(lengthOfLastWord("a"))
}
