package ui

import (
	"fmt"
)

func GetUsersData() (string, int, error) {
	var login string
	var password int

	printAuthInput()
	_, err := fmt.Scanln(&login, &password)

	return login, password, err
}

func GetMenuOption() (int, error) {
	var userOption int
	printMenu()
	askMenuOption()
	_, err := fmt.Scanln(&userOption)

	return userOption, err
}
