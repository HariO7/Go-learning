package main

import (
	"fmt"
)

func main() {
	var revenue, taxRate, expenses float64

	fmt.Print("What is the revenue receieved: ")
	fmt.Scan(&revenue)

	fmt.Print("What is the total expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("What is the tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate)
	ratio := earningsBeforeTax / profit

	fmt.Print("Earnings before tax: ")
	fmt.Println(earningsBeforeTax)

	fmt.Print("Total profit: ")
	fmt.Println(profit)

	fmt.Print("Ratio between the two: ")
	fmt.Println(ratio)
}
