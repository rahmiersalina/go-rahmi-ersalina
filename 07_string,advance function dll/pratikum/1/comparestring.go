package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compare("akashi", "aka"))
}
func compare(a, b string) string {
	if len(a) > len(b) {
		a = "long"
		b = "short"
	}
	if strings.Contains(a, b) {
		return a
	}
	return ""
}
