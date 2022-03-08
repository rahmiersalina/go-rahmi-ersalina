package main

import "fmt"

func main() {
	// IF,ELSE IF,ELSE
	hour := 20
	if hour < 12 {
		fmt.Println("selamat pagi")
	} else if hour < 18 {
		fmt.Println("selamat sore")
	} else {
		fmt.Println("selamat malam")
	}
}
