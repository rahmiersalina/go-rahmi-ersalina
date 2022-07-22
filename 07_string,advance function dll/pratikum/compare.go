package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	shortest, longest := a, b
	if len(a) > len(b) {
		longest = a
		shortest = b
	}
	if strings.Contains(longest, shortest) {
		return shortest
	}
	return ""
}

func main() {
	fmt.Print("Mari Kita Mencari Substring!")

	var a, b string
	fmt.Print("Enter A : ")
	fmt.Scanln(&a)

	fmt.Print("Enter B : ")
	fmt.Scanln(&b)

	fmt.Println(Compare(a, b))
}
