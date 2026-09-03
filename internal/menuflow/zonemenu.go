package menuflow

import (
	history "OStation/internal/eventhistory"
	"OStation/internal/orbital"
	"OStation/internal/orbital/storage"
	"OStation/internal/ui"
	"fmt"
)

func StorageMenu(orbitalStation *orbital.OrbitalStation, stationStorage *map[string]int, inventory *map[string]int, step *int, history *history.History) error {
	orbitalStation.ChangeZone(364)
	for {
		option, err := ui.GetStoregOption()
		err = ui.ValidateInput(err)

		if err != nil {
			return fmt.Errorf("Problem: %w", err)
		}

		if option == 3 {
			fmt.Println("You come back into controll room")
			orbitalStation.EnergyHandler()
			orbitalStation.ChangeZone(361)
			return nil
		}

		switch option {
		case 1:
			ui.PrintStorage(*stationStorage)
			orbitalStation.EnergyHandler()
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
			history.Add("Admin got some stuff from the storage")
			orbitalStation.EnergyHandler()
			orbitalStation.OxygenHandler(*step)
		}
		*step++
	}
}

func reactorMenu(station *orbital.OrbitalStation, inventory *map[string]int, step *int, history *history.History) error {
	for {
		option, err := ui.GetReactorMenu(*station.Zones[362])

		err = ui.ValidateInput(err)

		if err != nil {
			return fmt.Errorf("Problem: %w", err)
		}

		err = orbital.ZoneValidate(option, *station)

		if err != nil {
			return fmt.Errorf("Problem: %w", err)
		}

		if option == 3 {
			ui.LeftRoom()
			return nil
		}

		switch option {
		case 1:

			switch station.CurrentZone.Condition {
			case "unstable":
				ui.ReactorReport(*station.CurrentZone)

			case "stable":
				ui.ReactorStableReport(*station.CurrentZone)

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
			history.Add("Reactor stabilized")
			storage.SpecTool(inventory)
			ui.SuccesFixed(station.CurrentZone.Name)
			ui.SpesToolNotification()
			history.Add("Found a special tool")

		}

		station.OxygenHandler(*step)
		*step++

	}
}

func communicationMenu(station *orbital.OrbitalStation, inventory *map[string]int, step *int, history *history.History) error {
	for {
		option, err := ui.GetCommunicationMenu(*station.Zones[363])

		err = ui.ValidateInput(err)
		if err != nil {
			return fmt.Errorf("Problem: %w", err)
		}

		err = orbital.ZoneValidate(option, *station)

		if err != nil {
			return fmt.Errorf("ERROR: %w", err)
		}

		if option == 3 {
			ui.LeftRoom()
			return nil
		}

		switch option {
		case 1:
			switch station.CurrentZone.Condition {
			case "unstable":
				ui.CommunicationReport(*station.CurrentZone)

			case "stable":
				ui.CommunicationStable(*station.CurrentZone)
			}

		case 2:
			err = orbital.FixValidate(station.CurrentZone.StuffToFix, *inventory)

			if err != nil {
				return fmt.Errorf("ERROR: %w", err)
			}

			err = orbital.CommunicationValidate(*station)

			if err != nil {
				return fmt.Errorf("ERROR: %w", err)
			}

			storage.FixZone(&station.CurrentZone.StuffToFix, inventory)
			station.FixZone(363)
			history.Add("Communications restored")
			ui.SuccesFixed(station.CurrentZone.Name)
		}
		station.OxygenHandler(*step)
		*step++

	}
}

func lifeSupport(station *orbital.OrbitalStation, inventory *map[string]int, step *int, history *history.History) error {
	for {
		option, err := ui.GetLifeSupMenu(*station.CurrentZone)

		err = ui.ValidateInput(err)

		if err != nil {
			return fmt.Errorf("Problem: %w", err)
		}

		err = orbital.ZoneValidate(option, *station)

		if err != nil {
			return fmt.Errorf("ERROR: %w", err)
		}

		if option == 3 {
			ui.LeftRoom()
			return nil
		}

		switch option {
		case 1:
			switch station.CurrentZone.Condition {
			case "unstable":
				ui.LifeSupport(*station.CurrentZone)

			case "stable":
				ui.LifeStableSupport(*station.CurrentZone)
			}

		case 2:
			err = orbital.FixValidate(station.CurrentZone.StuffToFix, *inventory)

			if err != nil {
				return fmt.Errorf("ERROR: %w", err)
			}

			err = orbital.LifeSupport(*station)

			if err != nil {
				return fmt.Errorf("ERROR: %w", err)
			}

			storage.FixZone(&station.CurrentZone.StuffToFix, inventory)
			station.FixZone(365)
			ui.SuccesFixed(station.CurrentZone.Name)
			history.Add("Life Support fixed")
		}
		station.OxygenHandler(*step)
		*step++
	}
}
