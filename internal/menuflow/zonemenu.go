package menuflow

import (
	history "OStation/internal/eventhistory"
	state "OStation/internal/gamestate"
	"OStation/internal/orbital"
	"OStation/internal/orbital/storage"
	"OStation/internal/ui"
	"fmt"
)

func StorageMenu(station *orbital.Station, stationStorage map[string]int, inventory map[string]int, step *int, history *history.History) (error, bool) {
	station.ChangeZone(364)
	for {

		err := state.Handler(station)

		if err != nil {
			return err, true
		}
		option, err := ui.StoregOption()
		err = ui.ValidateInput(err)

		if err != nil {
			return fmt.Errorf("get menu option in the storage: %w", err), false
		}

		if option == 3 {
			fmt.Println("You come back into controll room")
			station.EnergyHandler()
			station.ChangeZone(361)
			return nil, false
		}

		switch option {
		case 1:
			ui.Storage(stationStorage)
			station.EnergyHandler()
		case 2:

			resource, amount, err := ui.TakingData()

			err = ui.ValidateInput(err)
			if err != nil {
				return fmt.Errorf("Problem: %w", err), false
			}

			err = storage.CheckData(resource, amount, stationStorage)

			if err != nil {
				return fmt.Errorf("get stuff from storage: %w", err), false

			}

			storage.TakeResource(stationStorage, inventory, resource, amount)
			storage.StorageCheck(stationStorage, resource)
			history.Add("Admin got some stuff from the storage")
			station.EnergyHandler()
			station.OxygenHandler(*step)
		}
		*step++
	}
}

func reactorMenu(station *orbital.Station, inventory map[string]int, step *int, history *history.History) (error, bool) {
	const reactorID int = 362
	for {
		err := state.Handler(station)

		if err != nil {
			return err, true
		}
		option, err := ui.ReactorMenu(*station.ZonesByID[reactorID])

		err = ui.ValidateInput(err)

		if err != nil {
			return fmt.Errorf("get option from reactor zone menu: %w", err), false
		}

		err = orbital.ZoneMenuValidate(option, *station)

		if err != nil {
			return fmt.Errorf("Problem: %w", err), false
		}

		if option == 3 {
			ui.LeftRoom()
			return nil, false
		}

		switch option {
		case 1:

			switch station.CurrentZone.Condition {
			case "unstable":
				ui.ReactorReport(*station.CurrentZone)

			case "stable":
				ui.ReactorStableReport(*station.CurrentZone)
			}

		case 2:

			err = orbital.FixValidate(station.CurrentZone.StuffToFix, inventory)

			if err != nil {
				return fmt.Errorf("can not fix reactor: %w", err), false
			}

			storage.FixZone(station.ZonesByID[362].StuffToFix, inventory)

			station.FixZone(reactorID)
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

func communicationMenu(station *orbital.Station, inventory map[string]int, step *int, history *history.History) (error, bool) {
	const communicationID int = 363
	for {

		err := state.Handler(station)
		if err != nil {
			return err, true
		}

		option, err := ui.CommunicationMenu(*station.ZonesByID[communicationID])

		err = ui.ValidateInput(err)
		if err != nil {
			return fmt.Errorf("get option from communication menu: %w", err), false
		}

		err = orbital.ZoneMenuValidate(option, *station)

		if err != nil {
			return fmt.Errorf("choose menu option: %w", err), false
		}

		if option == 3 {
			ui.LeftRoom()
			return nil, false
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
			err = orbital.FixValidate(station.CurrentZone.StuffToFix, inventory)

			if err != nil {
				return fmt.Errorf("can not fix zone: %w", err), false
			}

			err = orbital.CommunicationFixValidate(*station)

			if err != nil {
				return fmt.Errorf("can not fix communication: %w", err), false
			}

			storage.FixZone(station.CurrentZone.StuffToFix, inventory)
			station.FixZone(communicationID)
			history.Add("Communications restored")
			ui.SuccesFixed(station.CurrentZone.Name)
		}
		station.OxygenHandler(*step)
		*step++

	}
}

func lifeSupport(station *orbital.Station, inventory map[string]int, step *int, history *history.History) (error, bool) {
	const lifeSupportID int = 365
	for {

		err := state.Handler(station)

		if err != nil {
			return err, true
		}
		option, err := ui.LifeSupMenu(*station.CurrentZone)

		err = ui.ValidateInput(err)

		if err != nil {
			return fmt.Errorf("ger option from lifeSupport menu: %w", err), false
		}

		err = orbital.ZoneMenuValidate(option, *station)

		if err != nil {
			return fmt.Errorf("choose zone menu option: %w", err), false
		}

		if option == 3 {
			ui.LeftRoom()
			return nil, false
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
			err = orbital.FixValidate(station.CurrentZone.StuffToFix, inventory)

			if err != nil {
				return fmt.Errorf("can not fix life support: %w", err), false
			}

			err = orbital.LifeSupportFixValidate(*station)

			if err != nil {
				return fmt.Errorf("can not fix life support: %w", err), false
			}

			storage.FixZone(station.CurrentZone.StuffToFix, inventory)
			station.FixZone(lifeSupportID)
			ui.SuccesFixed(station.CurrentZone.Name)
			history.Add("Life Support fixed")
		}
		station.OxygenHandler(*step)
		*step++
	}
}
