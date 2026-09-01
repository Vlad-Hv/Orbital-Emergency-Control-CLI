package main

import (
	"OStation/internal/authorization"
	"OStation/internal/menuflow"
	"OStation/internal/orbital"
	"OStation/internal/orbital/storage"
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
	stationStorage := storage.CreateStorage()
	inventory := storage.CreateInventory()
	orbitalStation := orbital.CreateStation()

	for {
		option, err := ui.GetMenuOption()
		err = ui.ValidateInput(err)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 0 {
			fmt.Println("developer exit")
			break
		}

		switch option {
		case 1:
			fmt.Println(orbitalStation.CurrentZone)

		case 2:
			for {
				err := menuflow.ChooseZoneMenu()
				if err != nil {
					fmt.Println(err)
					continue
				}

				break
			}
		case 3:
			ui.AllZones(orbitalStation.Zones)
		case 4:
			for {
				err := menuflow.StorageMenu(&orbitalStation, &stationStorage, &inventory)

				if err != nil {
					fmt.Println(err)
					continue
				}

				break
			}
		case 5:

		case 6:
			ui.PrintInv(inventory)
		default:
			fmt.Println("\nMESSAGE: invalid option")
		}

	}

}
