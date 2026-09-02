package menuflow

import (
	"OStation/internal/orbital"
	"OStation/internal/ui"
	"fmt"
)

func ChooseZoneMenu(station *orbital.OrbitalStation, inventory *map[string]int, step *int) error {
	id, err := ui.GetRoomID()

	err = ui.ValidateInput(err)
	if err != nil {
		return fmt.Errorf("cannot change zone: %w", err)
	}

	switch id {
	case 361:
		ui.ControlRoom()
		*step++
		return nil
	case 362:
		station.ChangeZone(362)

		for {
			err := reactorMenu(station, inventory, step)

			if err != nil {
				fmt.Println(err)
				continue
			}
			station.EnergyHandler()
			station.ChangeZone(361)
			return nil
		}
		//validation and after for loop and bussiness logik of this zone
		//меню реактора делать в файле зоне меню

	case 363:
		//validation and after for loop and bussiness logik of this zone

		return nil
	case 364:
		//добавить надпись экспортируемую "войди черезз мейн меню"

		return nil
	case 365:
		//validation and after for loop and bussiness logik of this zone

		return nil
	default:
		ui.InvalidRoom()
		return nil
	}
}
