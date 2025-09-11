package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationAmount = 2.5
	var investmentAmount, interestRate, year float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the interest rate: ")
	fmt.Scan(&interestRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&year)

	futureValue := investmentAmount * math.Pow(1+interestRate/100, year)

	inFutureValue := futureValue / math.Pow(1+inflationAmount/100, year)

	fmt.Println(futureValue)
	fmt.Println(inFutureValue)
}
