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

func GetRoomID() (int, error) {
	var id int
	printChooseRoom()
	_, err := fmt.Scanln(&id)
	return id, err
}

func GetStoregOption() (int, error) {
	var option int
	printStorageMenu()
	_, err := fmt.Scanln(&option)

	return option, err
}

func GetTakingData() (string, int, error) {
	var resourse string
	var amount int

	askStorageTake()
	_, err := fmt.Scanln(&resourse, &amount)

	return resourse, amount, err
}
