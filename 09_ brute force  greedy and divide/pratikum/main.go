package main

import "fmt"

func BinarySearch(arr []int, x int) int {
	var a, b = 0, len(arr) - 1
	for a <= b {
		var median = a + (b-a)/2
		if arr[median] == x {
			return median
		}
		if x > arr[median] {
			a = median + 1
		} else {
			b = median - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))
	fmt.Println(BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5))
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53))
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100))
}
