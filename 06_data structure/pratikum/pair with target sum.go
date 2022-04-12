package main

import "fmt"

func mai() {
	arr := []int{1, 2, 3, 4, 6}
	fmt.Println(pairsum(arr, 6))
}
func pairsum(arr []int, target int) (true []int) {
	hmap := make(map[int]int)
	for index, a := range arr {
		var b = target - a
		if _, yes := hmap[b]; yes {
			true = append(true, hmap[b], index)
			break
		}
		hmap[a] = index
	}
	return true
}
