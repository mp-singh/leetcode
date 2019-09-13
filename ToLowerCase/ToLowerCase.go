package main

import (
	"log"
	"strings"
)

func ToLowerCase(str string) string {
	var temp []string
	for _, i := range str {
		if i >= 65 && i < 97 {
			temp = append(temp, string(i+32))
		} else {
			temp = append(temp, string(i))
		}
	}
	return strings.Join(temp, "")
}

func main() {
	str := ToLowerCase("Hello")
	log.Println(str)
}