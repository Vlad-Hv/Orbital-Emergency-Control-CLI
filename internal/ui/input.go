package ui

import (
	"fmt"
)

func GetUsersData() (string, int, error) {
	var login string
	var password int

	printAuthInputLog()
	fmt.Scanln(&login)

	printAuthInputPassw()
	_, err := fmt.Scanln(&password)

	return login, password, err
}
