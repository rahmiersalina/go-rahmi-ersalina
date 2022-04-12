package main

import (
	"fmt"
)

func main() {
	f := "123451234"
	numberappear(f)
}
func numberappear(f string) []int {
	hmap := make(map[string]int)
	for a := 0; a < len(f); a++ {
		b := string(f[a])
		val, ok := hmap[string(f[a])]
		if ok {
			hmap[b] = val + 1
		} else {
			hmap[b] = 1
		}
	}
	var number []int
	for b, value := range hmap {
		if value == 1 {
			number = append(number)
			fmt.Println("number", b)
		}
	}
	return number
}
