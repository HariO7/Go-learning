package main

import (
	"fmt"

	"example.com/struct/user"
)

func main() {
	var userfirstName = getUserInput("Please enter your firstName:")
	var userlastName = getUserInput("Please enter your lastName:")
	var userbirthDate = getUserInput("Please enter your birtday in dd/mm/yy format:")

	var appUser, err = user.New(userfirstName, userlastName, userbirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()

}

func getUserInput(promtText string) string {
	fmt.Println(promtText)
	var value string
	fmt.Scanln(&value)
	return value
}
