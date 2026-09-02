package menuflow

import (
	"OStation/internal/orbital"
	"OStation/internal/orbital/storage"
	"OStation/internal/ui"
	"fmt"
)

func StorageMenu(orbitalStation *orbital.OrbitalStation, stationStorage *map[string]int, inventory *map[string]int) error {
	orbitalStation.ChangeZone(364)
	for {
		option, err := ui.GetStoregOption()
		err = ui.ValidateInput(err)

		if err != nil {
			return fmt.Errorf("Problem: %w", err)
		}

		if option == 3 {
			fmt.Println("You come back into controll room")
			orbitalStation.ChangeZone(361)
			return nil
		}

		switch option {
		case 1:
			ui.PrintStorage(*stationStorage)
		case 2:
			//сделать как раз таки обработчик инвентаря, сделать изменине стореджа и перекладывать в инвентарь + валидация
			resourse, amount, err := ui.GetTakingData()

			err = ui.ValidateInput(err)
			if err != nil {
				return fmt.Errorf("Problem: %w", err)

			}

			err = storage.CheckStoragesData(resourse, amount, *stationStorage)

			if err != nil {
				return fmt.Errorf("Problem: %w", err)

			}

			storage.TakeResourse(stationStorage, inventory, resourse, amount)
			storage.StorageCheck(stationStorage, resourse)
		}
	}
}

func reactorMenu(station *orbital.OrbitalStation, inventory *map[string]int) error {
	for {
		option, err := ui.GetReactorMenu(*station.Zones[362])

		err = ui.ValidateInput(err)

		if err != nil {
			return fmt.Errorf("Problem: %w", err)
		}

		err = orbital.ReactorValidate(option, *station)

		if err != nil {
			return fmt.Errorf("Problem: %w", err)
		}

		if option == 3 {
			fmt.Println("Come back to the controll room!")
			return nil
		}

		switch option {
		case 1:

			switch station.CurrentZone.Condition {
			case "unstable":
				ui.ReactorReport(station.CurrentZone)

			case "stable":
				ui.ReactorStableReport(station.CurrentZone)

			default:
				fmt.Println("Developer mudak, naychis pisat pravilno")
			}

		case 2:
			//add validation is he able to fix
			err = orbital.FixValidate(station.CurrentZone.StuffToFix, *inventory)

			if err != nil {
				return fmt.Errorf("ERROR: %w", err)
			}

			storage.FixZone(&station.Zones[362].StuffToFix, inventory)

			station.FixZone(362)
			ui.SuccesFixed(station.CurrentZone.Name)
		}

	}
}
