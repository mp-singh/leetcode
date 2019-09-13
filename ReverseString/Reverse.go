package main

import "log"

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[len(s)-1-i], s[i] = s[i], s[len(s)-1-i]
	}
	log.Println(string(s))
}

func main() {
	reverseString([]byte("hello"))
}
