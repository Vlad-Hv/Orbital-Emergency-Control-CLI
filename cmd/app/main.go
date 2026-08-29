package main

import (
	"OStation/internal/authorization"
	"OStation/internal/orbital"
	"OStation/internal/ui"
	"fmt"
)

func main() {
	ui.FirstWarning()
	err := authorization.Auth()

	if err != nil {
		fmt.Println(err)
		return
	}
	ui.Access()
	orbitalStation := orbital.CreateStation()

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
			fmt.Println(orbitalStation.CurrentZone)

		case 2:
			id, err := ui.GetRoomID()

			err = ui.ValidateInput(err)
			if err != nil {
				fmt.Println(err)
				continue
			}

			switch id {
			case 361:
				//validation is it possiblle to get into this room and after for loop and bussiness logik of this zone

			case 362:
				//validation and after for loop and bussiness logik of this zone

			case 363:
				//validation and after for loop and bussiness logik of this zone

			case 364:
				//validation and after for loop and bussiness logik of this zone

			case 365:
				//validation and after for loop and bussiness logik of this zone

			}
		case 3:

		case 4:

		case 5:

		default:
			fmt.Println("\nMESSAGE: invalid option")
		}

	}

}
