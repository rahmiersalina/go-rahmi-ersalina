package main

import "fmt"

func Frog(jumps []int) int {
	minimumPossible := map[int]int{}
	for index, value := range jumps {
		minimumPossible[value] = index
	}
	if len(jumps) == 2 {
			minimumPossible[i] = Frog(n-1) + Frog(n-2)
		} else {
			
		}
	}
	return mimimumPossible[n]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 10, 50}))
}
