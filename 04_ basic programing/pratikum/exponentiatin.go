package main

import (
	"fmt"
)

func main() {
	number := 5
	pangkat := 3
	fmt.Println("number =", number)
	fmt.Println("pangkat =", pangkat)
	exponent(number, pangkat)
}
func exponent(number int, pangkat int) {
	var f int
	hasil := 1
	for f = 1; f <= pangkat; f++ {
		hasil = hasil * number
	}
}
