package main

import (
	"fmt"
	"structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "teste123")
	admin.OutputUserDetail()
	admin.ClearUserName()
	admin.OutputUserDetail()

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetail()
	appUser.ClearUserName()
	appUser.OutputUserDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
