package menuflow

import (
	history "OStation/internal/eventhistory"
	state "OStation/internal/gamestate"
	"OStation/internal/orbital"
	"OStation/internal/ui"
	"fmt"
)

func ZoneMenu(station *orbital.Station, inventory map[string]int, step *int, history *history.History) error {
	const controlRoomID int = 361
	const reactorID int = 362
	const communicationID int = 363
	const storageID int = 364
	const lifeSupportID int = 365

	for {
		id, err := ui.RoomID()

		err = ui.ValidateInput(err)
		if err != nil {
			fmt.Printf("cannot change zone: %v\n", err)
			continue
		}

		switch id {
		case controlRoomID:
			ui.ControlRoom()
			*step++
			station.OxygenHandler(*step)
			err = state.Handler(station)

			if err != nil {
				return err
			}

			return nil

		case reactorID:
			station.ChangeZone(reactorID)
			history.Add("Entered Reactor")
			for {
				err, isGameEnd := reactorMenu(station, inventory, step, history)

				if err != nil && isGameEnd {
					return err
				}
				if err != nil {
					fmt.Println(err)
					continue
				}
				station.EnergyHandler()
				station.OxygenHandler(*step)
				station.ChangeZone(controlRoomID)
				history.Add("Left Reactor")
				return nil
			}

		case communicationID:
			station.ChangeZone(communicationID)
			history.Add("Entered communication room")

			for {
				err, isGameEnd := communicationMenu(station, inventory, step, history)

				if err != nil && isGameEnd {
					return err
				}
				if err != nil {
					fmt.Println(err)
					continue
				}

				station.EnergyHandler()
				station.OxygenHandler(*step)
				station.ChangeZone(controlRoomID)
				history.Add("Left communication room")
				return nil
			}

		case storageID:
			ui.EnterFromMenu()
			return nil

		case lifeSupportID:
			station.ChangeZone(lifeSupportID)
			history.Add("Entered Life Support zone")

			for {
				err, isGameEnd := lifeSupport(station, inventory, step, history)

				if err != nil && isGameEnd {
					return err
				}

				if err != nil {
					fmt.Println(err)
					continue
				}

				station.EnergyHandler()
				station.OxygenHandler(*step)
				station.ChangeZone(controlRoomID)
				history.Add("Left Life Support zone")

				return nil
			}

		default:
			ui.InvalidRoom()
			return nil
		}
	}
}
