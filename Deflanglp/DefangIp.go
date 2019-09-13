package main

import (
	"log"
	"strings"
)

func DefangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

func main() {
		str := DefangIPaddr("1.1.1.1")
		log.Printntln(str)
}