package main

import (
	"OStation/internal/authorization"
	history "OStation/internal/eventhistory"
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
	ui.Warning()
	err := authorization.Auth()

	if err != nil {
		fmt.Println(err)
		return
	}
	ui.Access()

	history := history.Create()
	stationStorage := storage.Create()
	inventory := storage.CreateInventory()
	station := orbital.Create()

	for {
		err := state.Handler(&station)
		if err != nil {
			fmt.Println(err)
			return
		}

		option, err = ui.MenuOption(step)
		err = ui.ValidateInput(err)

		if err != nil {
			fmt.Println(err)
			continue
		}

		storage.AccessCard(step, inventory)
		if option == 0 {
			fmt.Println("developer exit")
			break
		}

		switch option {
		case 1:
			ui.StationReport(station)
			step++
			station.EnergyHandler()
			station.OxygenHandler(step)

		case 2:
			for {
				err := menuflow.ZoneMenu(&station, inventory, &step, &history)
				if err != nil {
					fmt.Println(err)
					return
				}

				break
			}
		case 3:
			ui.AllZones(station.ZonesByID)

		case 4:
			for {
				err, isGameEnd := menuflow.StorageMenu(&station, stationStorage, inventory, &step, &history)

				if err != nil && isGameEnd {
					fmt.Println(err)
					return
				}
				if err != nil {
					fmt.Println(err)
					continue
				}

				break
			}
		case 5:
			ui.HistoryReport(history)
		case 6:
			err := orbital.Signal(station)

			if err != nil {
				fmt.Println(err)
				continue
			}

			station.SendSignalSOS()
			ui.SignalSucces()
			history.Add("Emergency signal sent")

		case 7:
			ui.PrintInv(inventory)
		default:
			fmt.Println("\nMESSAGE: invalid option")
		}

	}

}
