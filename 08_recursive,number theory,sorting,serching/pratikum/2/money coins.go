package main

import "fmt"

var MONEY_LIST = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}

func main() {
	fmt.Println("result", moneychange(26123))
}
func moneychange(amount int) (result []int) {
	fmt.Printf("initial money %v/n", amount)
	for _, row := range MONEY_LIST {
		div := amount / row
		fmt.Printf("fraction %v get as much %v/n", row, div)
		if div == 0 {
			continue
		}
		for i := 0; i < div; i++ {
			result = append(result, row)
		}
		amount = amount - (row * div)
		fmt.Printf("money left now %v\n", amount)
		if amount == 0 {
			break
		}
	}
	return result
}
