package main

import "fmt"

func main() {
	var row int = 5
	var r int
	for f := 1; f <= row; f++ {
		for r = row - f; r > 0; r-- {
			fmt.Println(" ")
		}
		for spase := 1; spase <= f; spase++ {
			fmt.Println(" * ")
		}
		fmt.Println("")
	}
}
