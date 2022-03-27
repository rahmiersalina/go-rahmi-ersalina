package main

import "fmt"

func main() {
	var kalimat string
	fmt.Println("kalimat :")
	fmt.Scan(&kalimat)

	kebalikan := ispolindrome(kalimat)
	if kebalikan == true {
		fmt.Println("polindrome")
	} else {
		fmt.Println("bukan polindrome")
	}
}
func ispolindrome(input string) bool {
	for f := 0; f < len(input)/2; f++ {
		if input[f] != input[len(input)-f-1] {
			return false
		}
	}
	return true
}
