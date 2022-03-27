package main

import "fmt"

func main() {
	var bilangan, a int
	fmt.Println("bilangan :")
	fmt.Scanln(&bilangan)
	fmt.Println("faktor bilangan ", bilangan, "adalah")
	for a = 1; a <= bilangan; a++ {
		if bilangan%a == 0 {
			fmt.Print(a, ",")
		}
	}
}
