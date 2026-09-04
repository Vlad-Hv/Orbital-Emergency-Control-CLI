package menuflow

import (
	history "OStation/internal/eventhistory"
	state "OStation/internal/gamestate"
	"OStation/internal/orbital"
	"OStation/internal/ui"
	"fmt"
)

func ChooseZoneMenu(station *orbital.OrbitalStation, inventory *map[string]int, step *int, history *history.History) error {
	for {
		id, err := ui.GetRoomID()

		err = ui.ValidateInput(err)
		if err != nil {
			fmt.Printf("cannot change zone: %v\n", err)
			continue
		}

		switch id {
		case 361:
			ui.ControlRoom()
			*step++
			station.OxygenHandler(*step)
			err = state.GameState(station)

			if err != nil {
				return err
			}

			return nil
		case 362:
			station.ChangeZone(362)
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
				station.ChangeZone(361)
				history.Add("Left Reactor")
				return nil
			}

		case 363:
			station.ChangeZone(363)
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
				station.ChangeZone(361)
				history.Add("Left communication room")
				return nil
			}

		case 364:
			ui.EnterFromMenu()
			return nil

		case 365:
			station.ChangeZone(365)
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
				station.ChangeZone(361)
				history.Add("Left Life Support zone")

				return nil
			}

		default:
			ui.InvalidRoom()
			return nil
		}
	}
}
