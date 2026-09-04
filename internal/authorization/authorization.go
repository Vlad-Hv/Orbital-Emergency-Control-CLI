package authorization

import (
	"OStation/internal/ui"
	"errors"
)

func Auth() error {
	ui.PrintAuth()
	login, password, err := getLoginData()
	if err != nil {
		return err
	}
	for i := 0; i < 3; i++ {
		userLogin, userPassword, err := ui.GetUsersData()
		if err != nil {
			ui.InvalidAuth()
			continue
		}

		if login == userLogin && password == userPassword {
			return nil
		}
		ui.InvalidAuth()
	}

	return errors.New("Invalid user, Access denied!")
}
