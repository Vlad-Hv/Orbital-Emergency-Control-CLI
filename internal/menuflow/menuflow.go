package menuflow

import (
	history "OStation/internal/eventhistory"
	"OStation/internal/orbital"
	"OStation/internal/ui"

	"fmt"
)

func ChooseZoneMenu(station *orbital.OrbitalStation, inventory *map[string]int, step *int, history *history.History) error {
	id, err := ui.GetRoomID()

	err = ui.ValidateInput(err)
	if err != nil {
		return fmt.Errorf("cannot change zone: %w", err)
	}

	switch id {
	case 361:
		ui.ControlRoom()
		*step++
		station.OxygenHandler(*step)
		return nil
	case 362:
		station.ChangeZone(362)
		history.Add("Entered Reactor")
		for {
			err := reactorMenu(station, inventory, step, history)

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
		//validation and after for loop and bussiness logik of this zone
		//меню реактора делать в файле зоне меню

	case 363:
		station.ChangeZone(363)
		history.Add("Entered communication room")

		for {
			err := communicationMenu(station, inventory, step, history)

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
		//validation and after for loop and bussiness logik of this zone

	case 364:
		ui.EnterFromMenu()
		return nil

	case 365:
		station.ChangeZone(365)
		history.Add("Entered Life Support zone")

		for {
			err := lifeSupport(station, inventory, step, history)

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
