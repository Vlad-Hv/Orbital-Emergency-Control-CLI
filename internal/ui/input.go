package ui

import (
	"OStation/internal/orbital"
	"fmt"
)

func GetUsersData() (string, int, error) {
	var login string
	var password int

	printAuthInput()
	_, err := fmt.Scanln(&login, &password)

	return login, password, err
}

func GetMenuOption(step int) (int, error) {
	var userOption int
	printMenu()
	if step == 4 {
		accessCard()
	}
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

func GetReactorMenu(reactor orbital.Zone) (int, error) {
	var option int
	reactorMenu(reactor)

	_, err := fmt.Scanln(&option)

	return option, err
	//короче, я засыпаю. на сегодня мне проанализировать еще раз зоны ответственности, чтобы мб принт меню не отвечало за то, что выводить. 2 создать меню реактора и вывод состояния в зависимости от состояния, 3 добавить возможность починить реактор если достаточно ресурсов. 4 сделать, чтобы энергия тратилась, от каждого действия И после починки реактора переставало тратиться
}

func GetCommunicationMenu(communication orbital.Zone) (int, error) {
	var option int
	communicationMenu(communication)

	_, err := fmt.Scanln(&option)
	return option, err
}

func GetLifeSupMenu(lifeSup orbital.Zone) (int, error) {
	var option int

	lifeSupportMenu(lifeSup)
	_, err := fmt.Scanln(&option)

	return option, err
}
