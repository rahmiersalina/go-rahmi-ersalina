package main

import "fmt"

func main() {
	var number int
	fmt.Println("Inputlah Deret Bilangan Prima yang Ingin Kamu Ketahui : ")
	fmt.Scan(&number)
	fmt.Println("Bilangan Prima pada Deret tersebut adalah")
	fmt.Println(primeX(number))
}

func primeX(number int) int {
	result, i := 0, 2
	for {
		if val, isFound := primeNumber(i); isFound {
			result++
			if result == number {
				return val
			}
		}
		i++
	}
}

func primeNumber(number int) (int, bool) {
	if number == 2 || number == 3 || number == 5 || number == 7 {
		return 0, false
	} else if number%2 == 0 || number%3 == 0 || number%5 == 0 || number%7 == 0 {
		return 0, false
	} else {
		return number, true
	}
	return number, true
}

// func primeNumber(n int) (int, bool) {
// 	if n <= 1 {
// 		return 0, false
// 	}
// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return 0, false
// 		}
// 	}
// 	return n, true
// }
