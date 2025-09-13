package main

import "fmt"

func main() {
	accountBalance := 1000.0
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
			fmt.Println("Your total balance: ", accountBalance)
		default:
			fmt.Println("Thank you for using our app")
			return
		}
	}
}
