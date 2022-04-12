package main

import (
	"fmt"
)

func main() {
	arrA := []string{"justin", "zayn", "nichol", "evan", "jade"}
	arrB := []string{"kendall", "gigi", "bella", "selena", "hayley"}
	fmt.Println(ArryMerge(arrA, arrB))
}
func ArryMerge(arrA, arrB []string) (combined []string) {
	hmap := make(map[string]bool)
	for _, valueA := range arrA {
		hmap[valueA] = true
		combined = append(combined, valueA)
	}
	for _, valueB := range arrB {
		_, isfound := hmap[valueB]
		if isfound == false {
			hmap[valueB] = true
			combined = append(combined, valueB)
		}
	}
	return combined

}
