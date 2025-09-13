package main

import (
	"fmt"
	"math"
	"os"
)

const fileName = "investment.txt"

func main() {
	var revenue, taxRate, expenses float64
	var err error

	revenue, err = getValuesFromUser("What is the revenue receieved: ")
	handleErrors(err)
	expenses, err = getValuesFromUser("What is the total expenses: ")
	handleErrors(err)
	taxRate, err = getValuesFromUser("What is the tax rate: ")
	handleErrors(err)

	var earningsBeforeTax, profit, ratio float64 = calculateValues(revenue, expenses, taxRate)

	fmt.Print("Earnings before tax: ")
	fmt.Println(earningsBeforeTax)

	fmt.Print("Total profit: ")
	fmt.Println(profit)

	fmt.Printf("Ratio between the two: %.2f", ratio)
}

func getValuesFromUser(text string) (float64, error) {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 || math.IsNaN(value) {
		return 1000, fmt.Errorf("invalid input")
	}
	return value, nil
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	stringValue := fmt.Sprintf("\nEBT: %.1f,\n profit: %.1f,\n ratio: %.1f\n", earningsBeforeTax, profit, ratio)
	storeValues(stringValue)
	return earningsBeforeTax, profit, ratio
}

func handleErrors(err error) {
	if err != nil {
		fmt.Println("An error occured")
		fmt.Println(err)
		return
	}
}

func storeValues(value string) {
	data, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic(err)
	}

	data.Write([]byte(value))

}
