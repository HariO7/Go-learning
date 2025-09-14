package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	var userfirstName = getUserInput("Please enter your firstName:")
	var userlastName = getUserInput("Please enter your lastName:")
	var userbirthDate = getUserInput("Please enter your birtday in dd/mm/yy format:")

	var appUser = user{
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userbirthDate,
		createdAt: time.Now(),
	}

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()

}

func getUserInput(promtText string) string {
	fmt.Println(promtText)
	var value string
	fmt.Scan(&value)
	return value
}
