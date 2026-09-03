package main

import (
	"OStation/internal/authorization"
	state "OStation/internal/gamestate"
	"OStation/internal/menuflow"
	"OStation/internal/orbital"
	"OStation/internal/orbital/storage"
	"OStation/internal/ui"
	"fmt"
)

func main() {
	var step int
	var option int
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
		err := state.GameState(&orbitalStation)
		if err != nil {
			fmt.Println(err)
			return
		}

		option, err = ui.GetMenuOption(step)
		err = ui.ValidateInput(err)

		if err != nil {
			fmt.Println(err)
			continue
		}

		storage.AccessCard(step, &inventory)
		if option == 0 {
			fmt.Println("developer exit")
			break
		}

		switch option {
		case 1:
			ui.StationReport(orbitalStation)
			step++
			orbitalStation.EnergyHandler()
			orbitalStation.OxygenHandler(step)

		case 2:
			for {
				err := menuflow.ChooseZoneMenu(&orbitalStation, &inventory, &step)
				if err != nil {
					fmt.Println(err)
					continue
				}

				break
			}
		case 3:
			ui.AllZones(orbitalStation.Zones)
			/*step++
			orbitalStation.EnergyHandler()
			orbitalStation.OxygenHandler(step)*/
		case 4:
			for {
				err := menuflow.StorageMenu(&orbitalStation, &stationStorage, &inventory, &step)

				if err != nil {
					fmt.Println(err)
					continue
				}

				break
			}
		case 5:

		case 6:
			err := orbital.Signal(orbitalStation)

			if err != nil {
				fmt.Println(err)
				continue
			}

			orbitalStation.SendEmergySignal()
			ui.SignalSucces()

		case 7:
			ui.PrintInv(inventory)
		default:
			fmt.Println("\nMESSAGE: invalid option")
		}

	}

}
