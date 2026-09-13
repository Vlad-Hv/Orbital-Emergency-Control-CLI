package ui

import (
	"OStation/internal/orbital"
	"fmt"
)

func UsersData() (string, int, error) {
	var login string
	var password int

	authInput()
	_, err := fmt.Scanln(&login, &password)

	return login, password, err
}

func MenuOption(step int) (int, error) {
	var userOption int
	menu()
	if step == 4 {
		accessCard()
	}
	askMenuOption()
	_, err := fmt.Scanln(&userOption)

	return userOption, err
}

func RoomID() (int, error) {
	var id int
	chooseRoom()
	_, err := fmt.Scanln(&id)
	return id, err
}

func StoregOption() (int, error) {
	var option int
	storageMenu()
	_, err := fmt.Scanln(&option)

	return option, err
}

func TakingData() (string, int, error) {
	var resource string
	var amount int

	askStorageTake()
	_, err := fmt.Scanln(&resource, &amount)

	return resource, amount, err
}

func ReactorMenu(reactor orbital.Zone) (int, error) {
	var option int
	reactorMenu(reactor)

	_, err := fmt.Scanln(&option)

	return option, err
}

func CommunicationMenu(communication orbital.Zone) (int, error) {
	var option int
	communicationMenu(communication)

	_, err := fmt.Scanln(&option)
	return option, err
}

func LifeSupMenu(lifeSup orbital.Zone) (int, error) {
	var option int

	lifeSupportMenu(lifeSup)
	_, err := fmt.Scanln(&option)

	return option, err
}
