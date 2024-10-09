package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var ratio float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (earningsBeforeTax * taxRate)
	ratio = earningsBeforeTax / profit

	fmt.Printf("earning before tax: %v\n", earningsBeforeTax)
	fmt.Printf("profit: %v\n", profit)
  fmt.Printf("ratio: %v", ratio)

}
