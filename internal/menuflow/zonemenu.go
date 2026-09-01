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
