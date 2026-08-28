package main

import (
	"OStation/internal/authorization"
	"OStation/internal/ui"
	"fmt"
)

func main() {
	err := authorization.Auth()

	if err != nil {
		fmt.Println(err)
		return
	}
	ui.Access()

	for {
		option, err := ui.GetMenuOption()
		err = ui.ValidateInput(err)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 6 {
			fmt.Println("developer exit")
			break
		}

		switch option {
		case 1:

		case 2:

		case 3:

		case 4:

		case 5:

		default:
			fmt.Println("\nMESSAGE: invalid option")
		}

	}

}
