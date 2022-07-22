package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("knight fall")
	}
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	MinimumHeightOfKnight := 0
	pointTemp := 0
	for i := 0; i < len(dragonHead); i++ {
		for j := pointTemp; j < len(knightHeight); j++ {
			if knightHeight[j] >= dragonHead[i] {
				MinimumHeightOfKnight += knightHeight[j]
				pointTemp = j + 1
				break
			}
		}
	}
	if MinimumHeightOfKnight == 0 {
		fmt.Println("knight fall!")
	} else {
		fmt.Println(MinimumHeightOfKnight)
	}
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
	DragonOfLoowater([]int{5, 10}, []int{5})
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})
}
