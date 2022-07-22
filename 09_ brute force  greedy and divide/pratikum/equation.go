package main

import "fmt"

func SimpleEquations(a, b, c int) {
	for x := 1; x <= a; x++ {
		for y := x; y <= a; y++ {
			if b%(x*y) != 0 {
				continue
			}
			z := a - x - y
			if z <= 0 || z >= a {
				continue
			}
			if x*x+y*y+z*z == c {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
