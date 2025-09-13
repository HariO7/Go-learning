package main

import (
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func writeBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func readBalance() (float64, error) {
	balance, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, err
	}
	balanceValue, err := strconv.ParseFloat(string(balance), 64)

	if err != nil {
		return 1000, err
	}

	return balanceValue, nil
}

func main() {
	accountBalance, err := readBalance()

	if err != nil {
		fmt.Println("An Error occured")
		fmt.Println(err)
		fmt.Println("----------------------")
		panic(err)
	}
	fmt.Println("Welcome to Go BanK ")

	for {
		fmt.Println(`What do you want to do?
		1. Check balance
		2. Deposit Money
		3. Withdraw Money
		4. Exit`)

		var choice int
		fmt.Println("Please choose an option:")
		fmt.Scan(&choice)
		fmt.Println("Your choice:", choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is :", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Println("Deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				break
			}
			accountBalance += depositAmount
			writeBalance(accountBalance)
			fmt.Println("Your total balance :", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Println("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount")
				break
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient Balance")
				break
			}
			accountBalance -= withdrawAmount
			writeBalance(accountBalance)
			fmt.Println("Your total balance: ", accountBalance)
		default:
			fmt.Println("Thank you for using our app")
			return
		}
	}
}
