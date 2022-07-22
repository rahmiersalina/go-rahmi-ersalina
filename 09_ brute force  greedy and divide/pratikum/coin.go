package main

import "fmt"

var fraction = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}

func moneyCoins(money int) (MoneyCuts []int) {
	fmt.Printf("First Money %v\n", money)
	for _, row := range fraction {
		div := money / row
		if div == 0 {
			continue
		}
		for i := 0; i < div; i++ {
			MoneyCuts = append(MoneyCuts, row)
		}
		money = money - (row * div)
		if money == 0 {
			break
		}
	}
	return MoneyCuts
}

func main() {
	fmt.Println("MoneyCuts", moneyCoins(123))
	fmt.Println("MoneyCuts", moneyCoins(432))
	fmt.Println("MoneyCuts", moneyCoins(543))
	fmt.Println("MoneyCuts", moneyCoins(7752))
	fmt.Println("MoneyCuts", moneyCoins(15321))
}
