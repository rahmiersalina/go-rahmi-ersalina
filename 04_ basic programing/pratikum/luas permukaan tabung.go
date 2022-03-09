package main

import "fmt"

const phi = 3.14

func main() {

	// L = (2* phi * r * (r * r)) + (2 * phi * r * t)
	// L = 2 * phi * r * (r + t)
	T := 20.0
	r := 4.0
	L := 2 * phi * r * (r + T)

	fmt.Println(L)
}
