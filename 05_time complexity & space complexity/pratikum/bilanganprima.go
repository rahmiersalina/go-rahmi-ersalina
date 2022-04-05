package main

import "fmt"

func main() {
	n := 7
	fmt.Println("number", n)
	if primenumber(n) {
		fmt.Println("prime number")

	} else {
		fmt.Println("not primenumber")
	}

}
func primenumber(a int) bool {
	if a <= 1 {
		return false
	}
	for b := 2; b < a; b++ {
		sisa := a % b
		if sisa == 0 {
			return false
		}
	}
	return false
}
