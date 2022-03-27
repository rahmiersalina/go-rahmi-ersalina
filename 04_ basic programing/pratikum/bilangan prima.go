package main

import "fmt"

func main() {
	f := 10
	fmt.Println("number", f)
	if bilanganprima(f) {
		fmt.Println("bilangan prima")
	} else {

		fmt.Println("bukan bilangan prima")
	}
}

func bilanganprima(r int) bool {
	if r <= 1 {
		return false
	}
	for b := 2; b < r; b++ {
		tinggal := r % b
		if tinggal == 0 {
			return false
		}
	}
	return true
}
