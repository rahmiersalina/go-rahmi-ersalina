package main

import "fmt"

func fibo(n int) int {
	var dataBase map[int]int = map[int]int{}
	nilai, sudahTerhitung := dataBase[n]
	if sudahTerhitung {
		return nilai
	}
	if n <= 1 {
		dataBase[n] = n
	} else {
		dataBase[n] = fibo(n-1) + fibo(n-2)
	}
	return dataBase[n]
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}
